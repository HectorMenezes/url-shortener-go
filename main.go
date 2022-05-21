package main

import (
	"log"
	"net/http"

	url "github.com/HectorMenezes/url-shortener-go/controllers"
	db "github.com/HectorMenezes/url-shortener-go/db"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("Starting server...")

	db.Init()

	router := gin.Default()
	shortener := router.Group("/shortener", func(c *gin.Context) {
		auth := c.Request.Header.Get("Authorization")
		if auth == "" {
			c.String(http.StatusForbidden, "No Authorization header provided")
			c.Abort()
			return
		}
	})
	shortener.GET("/", url.GetUrls)
	shortener.POST("/", url.CreateUrl)
	shortener.GET("/:urlId", url.GetUrl)

	healthCheck := router.Group("/healthcheck")
	healthCheck.GET("/", url.GetHealthCheck)

	router.Run()
}
