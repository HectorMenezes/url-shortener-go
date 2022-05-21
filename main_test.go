package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	url "github.com/HectorMenezes/url-shortener-go/controllers"
	db "github.com/HectorMenezes/url-shortener-go/db"
	"github.com/HectorMenezes/url-shortener-go/models"
	"github.com/HectorMenezes/url-shortener-go/utils"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetupTestRoutes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	routes := gin.Default()
	return routes
}

func MockUrlCreate() string {
	url := models.Url{Url: "https://chess.com"}

	db.GetDB().Create(&url)
	return url.ID
}

func MockUrlDelete(id string) {
	var url models.Url
	db.GetDB().Delete(&url, id)
}

func TestGetAllUrls(t *testing.T) {

	db.Connect()
	idMockedUrl := MockUrlCreate()
	defer MockUrlDelete(idMockedUrl)

	router := SetupTestRoutes()
	router.GET("/", url.GetUrls)
	req, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code, "Should be 200")

}

func TestGetOneUrl(t *testing.T) {
	db.Connect()
	idMockedUrl := MockUrlCreate()
	defer MockUrlDelete(idMockedUrl)

	router := SetupTestRoutes()
	router.GET("/:urlId", url.GetUrl)
	req, _ := http.NewRequest("GET", fmt.Sprintf("/%s", idMockedUrl), nil)

	response := httptest.NewRecorder()
	router.ServeHTTP(response, req)

	assert.Equal(t, 301, response.Code, "Should redirect")
	assert.Equal(t, true, strings.Contains(response.Body.String(), `<a href="https://chess.com">Moved Permanently</a>.`))
}

func TestCreateUrl(t *testing.T) {
	db.Connect()

	router := SetupTestRoutes()
	router.POST("/", url.CreateUrl)
	req, _ := http.NewRequest("POST", "/", strings.NewReader(`{"url": "https://chess.com"}`))
	response := httptest.NewRecorder()
	router.ServeHTTP(response, req)

	var url models.UlrResponsePost
	json.Unmarshal(response.Body.Bytes(), &url)

	assert.Equal(t, url.Link, fmt.Sprintf("%s/%s",
		utils.GetEnv("SHORTENER_BASE_URL", ""),
		"1534637",
	))

	assert.Equal(t, url.OriginalUrl, "https://chess.com")
	MockUrlDelete(url.Link[len(url.Link)-7:])
}
