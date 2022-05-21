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

func GetUrls(c *gin.Context) {
	var urls []models.Url

	db := db.GetDB()
	db.Find(&urls)

	c.IndentedJSON(http.StatusOK, urls)
}

func GetUrl(c *gin.Context) {
	urlId := c.Param("urlId")
	var url models.Url
	db := db.GetDB()

	if err := db.Where("id = ?", urlId).First(&url).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.Redirect(301, url.Url)
}

func CreateUrl(c *gin.Context) {
	var url models.Url
	db := db.GetDB()

	if err := c.BindJSON(&url); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	validate := validator.New()

	if err := validate.Struct(url); err != nil {
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
