package comms

import (
	"fmt"
	"github.com/pelletier/go-toml/v2"
	"os"
)

type Config struct {
	Database DatabaseConfig `toml:"database"`
	Server   ServerConfig   `toml:"server"`
	// Add other fields as needed
}

// DatabaseConfig represents the database configuration
type DatabaseConfig struct {
	Host         string `toml:"host"`
	Port         int    `toml:"port"`
	User         string `toml:"user"`
	Password     string `toml:"password"`
	Database     string `toml:"database"`
	MaxOpenConns int    `toml:"maxOpenConns"`
	MaxIdleConns int    `toml:"maxIdleConns"`
}

// ServerConfig represents the server configuration
type ServerConfig struct {
	Address string `toml:"address"`
	Port    int    `toml:"port"`
}

func ReadConfig(filename string) (Config, error) {
	var config Config

	// Open and read the TOML file
	file, err := os.Open(filename)
	if err != nil {
		return config, fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	// Decode the TOML file into the Config struct
	if err := toml.NewDecoder(file).Decode(&config); err != nil {
		return config, fmt.Errorf("failed to decode TOML: %v", err)
	}

	return config, nil
}
