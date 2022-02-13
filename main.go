package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"hfs_backend/modules/product/productmodel"
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

	// CRUD

	products := r.Group("/products")
	{
		products.POST("", func(c *gin.Context) {

			var data productmodel.Product

			if err := c.ShouldBind(&data); err != nil {
				c.JSON(http.StatusUnauthorized, map[string]interface{}{
					"error": err.Error(),
				})
			}

			if err := db.Create(&data).Error; err != nil {
				c.JSON(http.StatusUnauthorized, map[string]interface{}{
					"error": err.Error(),
				})
			}
			c.JSON(http.StatusOK, data)

		})

		products.GET("/:id", func(c *gin.Context) {
			id, err := strconv.Atoi(c.Param("id"))

			if err != nil {
				c.JSON(http.StatusUnauthorized, map[string]interface{}{
					"error": err.Error(),
				})
				return
			}
			var data productmodel.Product

			if err := db.Where("id = ?", id).First(&data).Error; err != nil {
				c.JSON(http.StatusUnauthorized, map[string]interface{}{
					"error": err.Error(),
				})
			}
			c.JSON(http.StatusOK, data)
		})

		products.GET("", func(c *gin.Context) {

			var data []productmodel.Product

			var filter productmodel.ProductFilter

			c.ShouldBind(&filter)

			newDb := db

			if filter.Id > 0 {
				newDb.Where("id = ?", filter.Id)
			}

			if err := newDb.First(&data).Error; err != nil {
				c.JSON(http.StatusUnauthorized, map[string]interface{}{
					"error": err.Error(),
				})
			}
			c.JSON(http.StatusOK, data)
		})

		products.PATCH("/:id", func(c *gin.Context) {
			id, err := strconv.Atoi(c.Param("id"))

			if err != nil {
				c.JSON(401, map[string]interface{}{
					"error": err.Error(),
				})
				return
			}

			var data productmodel.Product

			if err := c.ShouldBind(&data); err != nil {
				c.JSON(401, map[string]interface{}{
					"error": err.Error(),
				})
				return
			}

			if err := db.Where("id = ?", id).Updates(&data).Error; err != nil {
				c.JSON(401, map[string]interface{}{
					"error": err.Error(),
				})

				return
			}

			c.JSON(http.StatusOK, gin.H{"ok": 1})
		})

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
