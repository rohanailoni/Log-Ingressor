package Models

import "time"

type UserAccess struct {
	User     string `bson:"user"`
	Password string `bson:"password"`
	Access   Access `bson:"Access"`
}

// Access represents the structure of the "Access" field in the MongoDB document
type Access struct {
	Regex    []string `bson:"regex"`
	Regular  []string `bson:"regular"`
	Wildcard []string `bson:"wildcard"`
}

// Config represents the structure of the configuration to be stored in the JSON file
type Config struct {
	Username string    `json:"username"`
	Time     time.Time `json:"time"`
	Access   Access    `json:"access"`
}
type FlagUser struct {
	User     string
	Password string
}
