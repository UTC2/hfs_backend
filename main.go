package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"hfs_backend/component"
	"hfs_backend/component/uploadprovider"
	"hfs_backend/middleware"
	"hfs_backend/modules/product/producttransport/ginproduct"
	"hfs_backend/modules/upload/uploadtransport/ginupload"
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

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	s3Provider := uploadprovider.NewS3Provider(s3bucketname, s3Region, s3Apikey, s3SecretKey, s3Domain)
	if err != nil {
		log.Fatalln(err)
	}

	if err := runService(db, s3Provider); err != nil {
		log.Fatalln(err)
	}
}

func runService(db *gorm.DB, provider uploadprovider.UploadProvider) error {

	r := gin.Default()
	appCtx := component.NewAppContext(db, provider)
	r.Use(middleware.Recover(appCtx))
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	// CRUD
	r.POST("/upload", ginupload.Upload(appCtx))

	products := r.Group("/products")
	{
		products.POST("", ginproduct.CreateProduct(appCtx))
		products.GET("/:id", ginproduct.GetProduct(appCtx))
		products.GET("", ginproduct.ListProduct(appCtx))
		products.PATCH("/:id", ginproduct.UpdateProduct(appCtx))
		products.DELETE("/:id", ginproduct.DeleteProduct(appCtx))
	}
	return r.Run()
}
