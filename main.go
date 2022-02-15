package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"hfs_backend/component"
	"hfs_backend/modules/product/productmodel"
	"hfs_backend/modules/product/producttransport/ginproduct"
	_ "hfs_backend/modules/product/producttransport/ginproduct"
	"log"
	"net/http"
	"os"
	"strconv"
	_ "time"
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

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	appCtx := component.NewAppContext(db)
	// CRUD
	products := r.Group("/products")
	{
		products.POST("", ginproduct.CreateProduct(appCtx))
		products.GET("/:id", ginproduct.GetProduct(appCtx))
		products.GET("", ginproduct.ListProduct(appCtx))
		products.PATCH("/:id", ginproduct.UpdateProduct(appCtx))

		products.DELETE("/:id", func(c *gin.Context) {

			id, err := strconv.Atoi(c.Param("id"))

			if err != nil {
				c.JSON(401, map[string]interface{}{
					"error": err.Error(),
				})
				return
			}

			if err := db.Table(productmodel.Product{}.TableName()).
				Where("id = ?", id).
				Delete(nil).Error; err != nil {
				c.JSON(401, map[string]interface{}{
					"error": err.Error(),
				})
				return
			}
			c.JSON(http.StatusOK, gin.H{"ok": 1})
		})
	}
	return r.Run()
}
