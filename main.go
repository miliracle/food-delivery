package main

import (
	"fooddelivery/common"
	"fooddelivery/component/appctx"
	"fooddelivery/component/uploadprovider"
	"fooddelivery/middleware"
	"fooddelivery/module/image/imagetransport/ginimage"
	"fooddelivery/module/restaurant/restauranttransport/ginrestaurent"
	"fooddelivery/module/user/userstorage"
	"fooddelivery/module/user/usertransport/ginuser"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dbConStr := common.CONFIG.DB_URI
	db, err := gorm.Open(mysql.Open(dbConStr), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	GCPCentificateFilePath := common.CONFIG.GOOGLE_APPLICATION_CREDENTIALS
	CloudStorageBucketName := common.CONFIG.GOOGLE_CLOUD_STORAGE_BUCKET_NAME

	uploadProvider, err := uploadprovider.NewGCPCloudStorageProvider(CloudStorageBucketName, GCPCentificateFilePath)

	if err != nil {
		log.Fatalln(err)
	}

	appCtx := appctx.NewAppContext(db, uploadProvider, common.CONFIG.SECRET_KEY)

	if err := runService(appCtx); err != nil {
		log.Fatalln(err)
	}
}

func runService(appCtx appctx.AppContext) error {
	userStore := userstorage.NewSQLStore(appCtx.GetMainDBConnection())
	r := gin.Default()
	r.Use(middleware.Recover(appCtx))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	v1 := r.Group("/v1")

	v1.POST("/register", ginuser.Register(appCtx))
	v1.POST("/login", ginuser.Login(appCtx))
	v1.GET("/profile", middleware.RequiredAuth(appCtx, userStore), ginuser.GetProfile(appCtx))

	images := v1.Group("/images")
	{
		images.GET("", ginimage.ListImage(appCtx))
		images.POST("", ginimage.UploadImage(appCtx))
		images.DELETE("/:id", ginimage.DeleteImage(appCtx))
	}

	restaurants := v1.Group("/restaurants", middleware.RequiredAuth(appCtx, userStore))
	{
		restaurants.GET("", ginrestaurent.ListRestaurant(appCtx))
		restaurants.POST("", ginrestaurent.CreateRestaurant(appCtx))
		restaurants.GET("/:id", ginrestaurent.GetRestaurant(appCtx))
		restaurants.PATCH("/:id", ginrestaurent.UpdateRestaurant(appCtx))
		restaurants.DELETE("/:id", ginrestaurent.DeleteRestaurant(appCtx))
	}

	return r.Run()
}
