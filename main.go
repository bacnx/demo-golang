package main

import (
	"fmt"
	"go-gin/model"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s",
		"localhost", "5432", "nxbac", "123456", "gin")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}

	sqlDB, err := db.DB()
	if err = sqlDB.Ping(); err != nil {
		log.Fatal("Failed to connect to the database")
	}

	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	router.POST("/items", func(c *gin.Context) {
		var item model.Item
		c.ShouldBind(&item)
		if err := db.Table("items").Create(&item).Error; err != nil {
			c.JSON(http.StatusFailedDependency, gin.H{
				"message": "Can't add to database",
			})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"message": "Create success",
		})
	})

	router.GET("/items", func(c *gin.Context) {
		var items []model.Item
		if err := db.Table("items").Find(&items).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Not found",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"posts": items,
		})
	})

	router.PUT("/items/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Params.ByName("id"))
		if err != nil {
			c.JSON(http.StatusFailedDependency, gin.H{
				"error": err.Error(),
			})
			return
		}

		var item model.Item
		if err := db.Where("id = ?", id).Table("items").First(&item).Error; err != nil {
			c.JSON(http.StatusFailedDependency, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.ShouldBind(&item)
		db.Save(&item)
		c.JSON(http.StatusOK, gin.H{
			"data": item,
		})
	})

	router.DELETE("items/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Params.ByName("id"))
		if err != nil {
			c.JSON(http.StatusFailedDependency, gin.H{
				"error": err.Error(),
			})
			return
		}

		var item model.Item
		if err := db.Where("id = ?", id).Table("items").Delete(&item).Error; err != nil {
			c.JSON(http.StatusFailedDependency, gin.H{
				"message": "Can't find with id",
				"error":   err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"data": item,
		})
	})

	router.Run("localhost:8000")
}
