package common

import (
	"log"
	"strconv"

	"github.com/spf13/viper"
)

type Config struct {
	MYSQL_VERSION                    string
	DB_URI                           string
	VOLUME_PATH                      string
	GOOGLE_APPLICATION_CREDENTIALS   string
	GOOGLE_CLOUD_STORAGE_BUCKET_NAME string
	GOOGLE_CDN                       string
	SECRET_KEY                       string
	ACCESS_TOKEN_EXPIRE_TIME         int
	REFRESH_TOKEN_EXPIRE_TIME        int
}

var CONFIG Config = Config{
	MYSQL_VERSION:                    viperEnvVariable("MYSQL_VERSION"),
	DB_URI:                           viperEnvVariable("DB_URI"),
	VOLUME_PATH:                      viperEnvVariable("VOLUME_PATH"),
	GOOGLE_APPLICATION_CREDENTIALS:   viperEnvVariable("GOOGLE_APPLICATION_CREDENTIALS"),
	GOOGLE_CLOUD_STORAGE_BUCKET_NAME: viperEnvVariable("GOOGLE_CLOUD_STORAGE_BUCKET_NAME"),
	GOOGLE_CDN:                       viperEnvVariable("GOOGLE_CDN"),
	SECRET_KEY:                       viperEnvVariable("SECRET_KEY"),
	ACCESS_TOKEN_EXPIRE_TIME:         viperEnvVariableInt("ACCESS_TOKEN_EXPIRE_TIME"),
	REFRESH_TOKEN_EXPIRE_TIME:        viperEnvVariableInt("REFRESH_TOKEN_EXPIRE_TIME_KEY"),
}

func viperEnvVariable(key string) string {

	// SetConfigFile explicitly defines the path, name and extension of the config file.
	// Viper will use this and not check any of the config paths.
	// .env - It will search for the .env file in the current directory
	viper.SetConfigFile(".env")

	// Find and read the config file
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	// viper.Get() returns an empty interface{}
	// to get the underlying type of the key,
	// we have to do the type assertion, we know the underlying value is string
	// if we type assert to other type it will throw an error
	value, ok := viper.Get(key).(string)

	// If the type is a string then ok will be true
	// ok will make sure the program not break
	if !ok {
		log.Fatalf("%s: Invalid type assertion", key)
	}

	return value
}

func viperEnvVariableInt(key string) int {
	value := viperEnvVariable(key)

	i, err := strconv.Atoi(value)
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)

	}
	return i
}
