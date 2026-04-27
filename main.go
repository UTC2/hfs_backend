package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"hfs_backend/component"
	"hfs_backend/component/uploadprovider"
	"hfs_backend/middleware"
  "hfs_backend/modules/house/housetransport/ginhouse"
  "hfs_backend/modules/product/producttransport/ginproduct"
	"hfs_backend/modules/upload/uploadtransport/ginupload"
  "hfs_backend/modules/user/usertransport/ginuser"
  "log"
	"net/http"
	"os"
)

func main() {

	dsn := os.Getenv("DBConnectionStr")
	s3bucketname := os.Getenv("S3BucketName")
	s3Region := os.Getenv("S3Region")
	s3Apikey := os.Getenv("S3APIKEY")
	s3SecretKey := os.Getenv("S3Secretkey")
	s3Domain := os.Getenv("S3Domain")
  secretKey := os.Getenv("SYSTEM_SECRET")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	s3Provider := uploadprovider.NewS3Provider(s3bucketname, s3Region, s3Apikey, s3SecretKey, s3Domain)
	if err != nil {
		log.Fatalln(err)
	}

	if err := runService(db, s3Provider,secretKey); err != nil {
		log.Fatalln(err)
	}
}

func runService(db *gorm.DB, provider uploadprovider.UploadProvider, secretkey string) error {
	r := gin.Default()
	appCtx := component.NewAppContext(db, provider, secretkey)
	r.Use(middleware.Recover(appCtx))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	v1 := r.Group("/v1")
	{
		// public auth endpoints
		v1.POST("/register", ginuser.Register(appCtx))
		v1.POST("/login", ginuser.Login(appCtx))
		v1.POST("/refresh", ginuser.Refresh(appCtx))

		// authenticated user endpoints
		v1.GET("/profile", middleware.RequiredAuth(appCtx), ginuser.GetProfile(appCtx))
		v1.POST("/upload", middleware.RequiredAuth(appCtx), ginupload.Upload(appCtx))

		// products: public read, auth'd write
		v1.GET("/products", ginproduct.ListProduct(appCtx))
		v1.GET("/products/:id", ginproduct.GetProduct(appCtx))
		v1.POST("/products", middleware.RequiredAuth(appCtx), ginproduct.CreateProduct(appCtx))
		v1.PATCH("/products/:id", middleware.RequiredAuth(appCtx), ginproduct.UpdateProduct(appCtx))
		v1.DELETE("/products/:id", middleware.RequiredAuth(appCtx), ginproduct.DeleteProduct(appCtx))

		// houses: public read, auth'd write
		v1.GET("/houses", ginhouse.ListHouse(appCtx))
		v1.GET("/houses/:id", ginhouse.GetHouse(appCtx))
		v1.POST("/houses", middleware.RequiredAuth(appCtx), ginhouse.CreateHouse(appCtx))
		v1.PATCH("/houses/:id", middleware.RequiredAuth(appCtx), ginhouse.UpdateHouse(appCtx))
		v1.DELETE("/houses/:id", middleware.RequiredAuth(appCtx), ginhouse.DeleteHouse(appCtx))
	}

	return r.Run()
}
