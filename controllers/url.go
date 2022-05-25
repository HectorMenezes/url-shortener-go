package url

import (
	"fmt"
	"net/http"

	"github.com/HectorMenezes/url-shortener-go/db"
	"github.com/HectorMenezes/url-shortener-go/models"
	"github.com/HectorMenezes/url-shortener-go/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// GetUrls return all urls from database.
func GetUrls(c *gin.Context) {
	var urls []models.Url

	db := db.GetDB()
	db.Find(&urls)

	c.IndentedJSON(http.StatusOK, urls)
}

// GetUrl redirect to a url based on hash id.
func GetUrl(c *gin.Context) {
	urlId := c.Param("urlId")
	var url models.Url
	db := db.GetDB()

	if err := db.Where("id = ?", urlId).First(&url).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "Url not found!",
		})
		return
	}
	c.Redirect(http.StatusTemporaryRedirect, url.Url)
}

// CreateUrl validates the payload and then inset into database.
func CreateUrl(c *gin.Context) {
	var url models.Url
	db := db.GetDB()

	if err := c.BindJSON(&url); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := validator.New().Struct(url); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	db.Create(&url)

	link := fmt.Sprintf("%s/%s",
		utils.GetEnv("SHORTENER_BASE_URL", ""),
		url.ID,
	)
	c.JSON(201, models.UlrResponsePost{Link: link, OriginalUrl: url.Url})
}
