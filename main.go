package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"hfs_backend/component"
	"hfs_backend/middleware"
	"hfs_backend/modules/product/producttransport/ginproduct"
	"hfs_backend/modules/upload/uploadtransport/ginupload"
	"log"
	"net/http"
	"os"
)

func main() {
	dsn := os.Getenv("DBConnectionStr")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	if err := runService(db); err != nil {
		log.Fatalln(err)
	}
}

func runService(db *gorm.DB) error {

	r := gin.Default()
	appCtx := component.NewAppContext(db)
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
