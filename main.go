package main

import (
	"fooddelivery/component"
	"fooddelivery/module/restaurant/restauranttransport/ginrestaurent"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dbConStr := viperEnvVariable("DB_URI")
	db, err := gorm.Open(mysql.Open(dbConStr), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	if err := runService(db); err != nil {
		log.Fatalln(err)
	}
}

func runService(db *gorm.DB) error {
	appCtx := component.NewAppContext(db)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	restaurants := r.Group("/restaurants")
	{
		restaurants.POST("", ginrestaurent.CreateRestaurant(appCtx))
	}

	return r.Run()
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
		log.Fatalf("Invalid type assertion")
	}

	return value
}
