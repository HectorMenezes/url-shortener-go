package main

import (
	"log"

	url "github.com/HectorMenezes/url-shortener-go/controllers"
	cache "github.com/HectorMenezes/url-shortener-go/cache"
	db "github.com/HectorMenezes/url-shortener-go/db"
    docs "github.com/HectorMenezes/url-shortener-go/docs"
    swaggerfiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"
    "github.com/jinzhu/gorm"

	"github.com/gin-gonic/gin"
)

// @title Url Shortener
// @version 1.0
// @description Simple URL shortener API
// @contact.name Hector Menezes
// @contact.email hector.menezes@gmail.com
// @license.name MIT License
// @license.url https://mit-license.org/
// @host localhost:8080
// @securityDefinitions.apiKey ApiKeyAuth 
// @in header 
// @name Authorization
func main() {
	log.Println("Starting server...")
	db.Connect(gorm.Open)
	db.Migrate()
    cache.Start()

	defer db.GetDB().Close()

	router := gin.Default()

    v1 := router.Group("/api/v1")

    docs.SwaggerInfo.BasePath = "/api/v1"

	shortener := v1.Group("/shortener", url.AuthMiddleware)
	shortener.GET("/", url.GetUrls)
	shortener.POST("/", url.CreateUrl)
	shortener.GET("/:urlId", url.GetUrl)

	healthCheck := v1.Group("/healthcheck")

	healthCheck.GET("/", url.GetHealthCheck([]url.ServiceChecker{
		{
			Name:          "database",
			HealthChecker: db.HealthcheckDB,
		},
	}))
    router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.Run()
}
