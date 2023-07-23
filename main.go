package main

import (
	"log"

	url "github.com/HectorMenezes/url-shortener-go/controllers"
	cache "github.com/HectorMenezes/url-shortener-go/cache"
	db "github.com/HectorMenezes/url-shortener-go/db"
	"github.com/jinzhu/gorm"

	"github.com/gin-gonic/gin"
)

// main connects to database, perform migrations and set
// all routes of app.
func main() {
	log.Println("Starting server...")
	db.Connect(gorm.Open)
	db.Migrate()
    cache.Start()

	defer db.GetDB().Close()

	router := gin.Default()

	shortener := router.Group("/shortener", url.AuthMiddleware)
	shortener.GET("/", url.GetUrls)
	shortener.POST("/", url.CreateUrl)
	shortener.GET("/:urlId", url.GetUrl)

	healthCheck := router.Group("/healthcheck")

	healthCheck.GET("/", url.GetHealthCheck([]url.ServiceChecker{
		{
			Name:          "database",
			HealthChecker: db.HealthcheckDB,
		},
	}))

	router.Run()
}
