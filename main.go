package main

import (
	"fooddelivery/common"
	"fooddelivery/component"
	"fooddelivery/component/uploadprovider"
	"fooddelivery/middleware"
	"fooddelivery/module/image/imagetransport/ginimage"
	"fooddelivery/module/restaurant/restauranttransport/ginrestaurent"
	"fooddelivery/module/user/usertransport/ginuser"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dbConStr := common.CONFIG.DB_URI
	db, err := gorm.Open(mysql.Open(dbConStr), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	GCPCentificateFilePath := viperEnvVariable("GOOGLE_APPLICATION_CREDENTIALS")
	CloudStorageBucketName := viperEnvVariable("GOOGLE_CLOUD_STORAGE_BUCKET_NAME")

	uploadProvider, err := uploadprovider.NewGCPCloudStorageProvider(CloudStorageBucketName, GCPCentificateFilePath)

	if err != nil {
		log.Fatalln(err)
	}

	if err := runService(db, uploadProvider); err != nil {
		log.Fatalln(err)
	}
}

func runService(db *gorm.DB, uploadProvider uploadprovider.UploadProvider) error {
	appCtx := component.NewAppContext(db, uploadProvider)

	r := gin.Default()
	r.Use(middleware.Recover(appCtx))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/register", ginuser.Register(appCtx))

	images := r.Group("/images")
	{
		images.GET("", ginimage.ListImage(appCtx))
		images.POST("", ginimage.UploadImage(appCtx))
		images.DELETE("/:id", ginimage.DeleteImage(appCtx))
	}

	restaurants := r.Group("/restaurants")
	{
		restaurants.GET("", ginrestaurent.ListRestaurant(appCtx))
		restaurants.POST("", ginrestaurent.CreateRestaurant(appCtx))
		restaurants.GET("/:id", ginrestaurent.GetRestaurant(appCtx))
		restaurants.PATCH("/:id", ginrestaurent.UpdateRestaurant(appCtx))
		restaurants.DELETE("/:id", ginrestaurent.DeleteRestaurant(appCtx))
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
