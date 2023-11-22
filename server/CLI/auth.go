package CLI

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/rohanailoni/Log-Ingressor/CLI/Models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func Authenticate(username string, password string) error {
	connectionString := fmt.Sprintf("mongodb+srv://rohan:%s@errorlevelshard.k2q89yg.mongodb.net/?retryWrites=true&w=majority", "cZksqzbBIpkRg1Gp")

	// Set up MongoDB client options
	clientOptions := options.Client().ApplyURI(connectionString)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Println(err)
		return err
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Println(err)
		return err
	}

	fmt.Println("Connected to MongoDB!")

	// Choose the database and collection
	database := client.Database("Logger")
	collection := database.Collection("Users")

	// Create a document to insert

	filter := bson.D{{"user", username}, {"password", password}}

	var result Models.UserAccess
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Println(err)
		return err
	}
	config := Models.Config{
		Username: result.User,
		Time:     time.Now(),
		Access:   result.Access,
	}
	configJSON, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		log.Println(err)
		return err
	}
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Println(err)
		return err
	}

	jsonFileName := "cli.json"
	AdditionConfigPath := "/.config/dyte/"
	jsonFilePath := filepath.Join(homeDir+AdditionConfigPath, jsonFileName)

	// Create the directory if it does not exist
	err = os.MkdirAll(filepath.Dir(jsonFilePath), 0755)
	if err != nil {
		log.Println(err)
		return err
	}

	// Write the JSON to the file
	err = os.WriteFile(jsonFilePath, configJSON, 0644)
	if err != nil {
		log.Println(err)
		return err
	}

	fmt.Printf("Configuration for user %s written to %s\n", username, jsonFilePath)
	return nil
}

func CheckAuthAndPermission(flagValue Models.Flagvalue) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Println(err)
		return err
	}

	jsonFileName := "cli.json"
	AdditionConfigPath := "/.config/dyte/"
	jsonFilePath := filepath.Join(homeDir+AdditionConfigPath, jsonFileName)
	file, err := os.Open(jsonFilePath)
	if err != nil {
		log.Println(err)
		return err
	}
	defer file.Close()

	// Decode the JSON data into a Config struct
	var config Models.Config
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		log.Println(err)
		return err
	}

	if flag := checkTimeExpiry(config.Time); !flag {
		return errors.New("timeout please authenticate again")
	}
	if !checkIfFlagsAreAuthenticated(flagValue, config) {
		return errors.New("user Not authenticated to use this flag check with admin")
	}
	return nil
}

/*
*This is quite a Direct Trick to find if ther fileds are present in struct.
 */

func checkIfFlagsAreAuthenticated(flagValue Models.Flagvalue, config Models.Config) bool {
	var regexAccessFields []string
	var regularAccessFields []string
	var wildcardAccessFields []string
	for _, value := range config.Access.Regular {
		regularAccessFields = append(regularAccessFields, strings.ToLower(value))
	}
	for _, value := range config.Access.Regex {
		regexAccessFields = append(regexAccessFields, strings.ToLower(value))
	}
	for _, value := range config.Access.Wildcard {
		wildcardAccessFields = append(wildcardAccessFields, strings.ToLower(value))
	}

	HasFieldPermission := true
	HasFieldPermission = HasFieldPermission && checkAvailability("level", regexAccessFields, regularAccessFields, wildcardAccessFields, flagValue.Level)
	HasFieldPermission = HasFieldPermission && checkAvailability("resourceId", regexAccessFields, regularAccessFields, wildcardAccessFields, flagValue.ResourceId)
	HasFieldPermission = HasFieldPermission && checkAvailability("spanid", regexAccessFields, regularAccessFields, wildcardAccessFields, flagValue.SpanId)
	HasFieldPermission = HasFieldPermission && checkAvailability("commit", regexAccessFields, regularAccessFields, wildcardAccessFields, flagValue.Commit)
	HasFieldPermission = HasFieldPermission && checkAvailability("message", regexAccessFields, regularAccessFields, wildcardAccessFields, flagValue.Message)
	HasFieldPermission = HasFieldPermission && checkAvailability("traceId", regexAccessFields, regularAccessFields, wildcardAccessFields, flagValue.TraceId)

	return HasFieldPermission
}
func checkAvailability(types string, regex, regular, wildcard []string, set Models.Valueset) bool {
	if set.HasRegular() && !contains(regular, types) {
		if !contains(regular, "all") {
			return false
		}

	}
	if set.HasRegex() && !contains(regex, types) {
		if !contains(regex, "all") {
			return false
		}
	}
	if set.HasWildcard() && !contains(wildcard, types) {
		if !contains(wildcard, "all") {
			return false
		}
	}
	return true
}

// return true if time is not expired
func checkTimeExpiry(timestamp time.Time) bool {

	timeDifference := time.Now().Sub(timestamp)

	return timeDifference.Hours() < 7*24
}

func contains(slice []string, element string) bool {
	for _, item := range slice {
		if strings.ToLower(item) == strings.ToLower(element) {
			return true
		}
	}
	return false
}
