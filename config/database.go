package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Store Database configuration parameters
type DatabaseConfiguration struct {
	DbType      string
	DbUsername  string
	DbPassword  string
	DbName      string
	DbHost      string
	DbPort      string
	DatabaseURL string
}

// Create a global variable to store database configurations
var DatabaseConfig DatabaseConfiguration

// Function to get configration related to database from configuration file
func GetDatabaseConfig() (err error) {
	// set configuration file for viper to process
	viper.SetConfigFile("config.yml")
	// Overwrite values from file
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	// Load the database configuration in the database struct
	DatabaseConfig.DbType = viper.GetString("DB_TYPE")
	DatabaseConfig.DbUsername = viper.GetString("DB_USERNAME")
	DatabaseConfig.DbPassword = viper.GetString("DB_PASSWORD")
	DatabaseConfig.DbName = viper.GetString("DB_NAME")
	DatabaseConfig.DbHost = viper.GetString("DB_HOST")
	DatabaseConfig.DbPort = viper.GetString("DB_PORT")
	err = generateConnectionURL()
	fmt.Println(DatabaseConfig)
	return err
}

// Generate connection string required for database connection
func generateConnectionURL() error {
	if DatabaseConfig.DbType == "postgres" {
		DatabaseConfig.DatabaseURL = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", DatabaseConfig.DbHost, DatabaseConfig.DbUsername, DatabaseConfig.DbPassword, DatabaseConfig.DbName, DatabaseConfig.DbPort)
		return nil
	}
	DatabaseConfig.DatabaseURL = ""
	return fmt.Errorf("Only Postgres Database supported. Invalid database: %s", DatabaseConfig.DbType)
}
