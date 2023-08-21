package url

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/HectorMenezes/url-shortener-go/cache"
	"github.com/HectorMenezes/url-shortener-go/db"
	"github.com/HectorMenezes/url-shortener-go/models"
	"github.com/HectorMenezes/url-shortener-go/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var BASE_URL = utils.GetEnv("SHORTENER_BASE_URL", "")

// @BasePath /api/v1
// @Summary Retrieve all urls
// @Schemes
// @Description Retrieve all urls registered in the system
// @Tags Shortener
// @Accept json
// @Produce json
// @Success 200 {array} models.UrlResponse "Ok"
// @Router /shortener/ [get]
// @Security ApiKeyAuth
func GetUrls(c *gin.Context) {
	var urls []models.Url

	db := db.GetDB()
	db.Find(&urls)

    var res []models.UrlResponse
    for _, url := range urls {
        link := fmt.Sprintf("%s/%s",
            BASE_URL,
            url.ID,
        )
        res = append(res, models.UrlResponse{
            Link: link, 
            OriginalUrl: url.Url, 
            CreatedAt: url.CreatedAt,
        })
    }
	c.IndentedJSON(http.StatusOK, res)
}

// @BasePath /api/v1
// @Summary Retrieve one URL
// @Schemes
// @Description Retrieve one URL based on ID
// @Tags Shortener
// @Accept json
// @Produce json
// @Param urlId path int true "Id of URL"
// @Success 307 {object} nil "Redirects to request URL"
// @Failure 404 
// @Failure 500
// @Router /shortener/{urlId} [get]
// @Security ApiKeyAuth
func GetUrl(c *gin.Context) {
	urlId := c.Param("urlId")

    cache := cache.GetCache()

    items, ok := cache.Read(urlId)
    var url models.Url

    if !ok {
        db := db.GetDB()

        if err := db.Where("id = ?", urlId).First(&url).Error; err != nil {
            c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
                "message": "Url not found!",
            })
            return
        }
        cache.Update(urlId, url)

    } else {
        err := json.Unmarshal(items, &url)
        if err != nil {
            c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
                "message": "Internal error!",
            })
            return
        }
    }
    c.Redirect(http.StatusTemporaryRedirect, url.Url)
}

// @BasePath /api/v1
// @Summary Short ulr
// @Schemes
// @Description Register a shorter url into system
// @Tags Shortener
// @Accept json
// @Produce json
// @Param url body models.UrlRequest true "Url information to short"
// @Success 201 {object} models.UrlResponse "Short version of URL created with success"
// @Router /shortener/ [post]
// @Security ApiKeyAuth
func CreateUrl(c *gin.Context) {
	var url models.UrlRequest
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
    databaseUrl := models.Url{Url: url.Url}

	db.Create(&databaseUrl)

	link := fmt.Sprintf("%s/%s",
		BASE_URL,
		databaseUrl.ID,
	)
	c.JSON(201, models.UrlResponse{Link: link, OriginalUrl: url.Url, CreatedAt: databaseUrl.CreatedAt})
}
