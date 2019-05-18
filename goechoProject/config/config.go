package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)


// Config stores the application-wide configurations
type appConfig struct {
	// the server port. Defaults to 8080
	ServerPort int `validate:"required"`
	// the data source name (DSN) for connecting to the database. required.
	DSN string `validate:"required"`
	// the DB_NAME for connecting to the database. required.
	DBNAME string `validate:"required"`
	// the signing method for JWT. Defaults to "HS256"
	JWTSigningMethod string `validate:"required"`
	// JWT signing key. required.
	JWTSigningKey string `validate:"required"`
	// JWT verification key. required.
	JWTVerificationKey string `validate:"required"`
}

type Config struct {
	App appConfig
}

// LoadConfig loads configuration from the given list of paths and populates it into the Config variable.
// The configuration file(s) should be named as app.yaml.
// Environment variables with the prefix "RESTFUL_" in their names are also read automatically.


// init is invoked before main()
func init(){
	// loads values from .env into the system
	if err := godotenv.Load("config.env"); err != nil {
		log.Print("No .env file found")
	}
}




func LoadConfig() *Config {
	os.Setenv("jwt_signing_method","HS256")
	return &Config{
		App: appConfig{
			ServerPort: getEnvAsInt("ServerPort", 8080),
			DSN: getEnv("dsn", ""),
			DBNAME:getEnv("DBNAME", ""),
			JWTSigningMethod:getEnv("jwt_signing_method",""),
			JWTSigningKey:getEnv("jwt_signing_key",""),
			JWTVerificationKey:getEnv("jwt_verification_key",""),
		},
	}
}

// Simple helper function to read an environment or return a default value
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}


func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultVal
}