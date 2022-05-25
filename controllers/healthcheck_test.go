package url

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/HectorMenezes/url-shortener-go/db"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func TestHealthcheck(t *testing.T) {
	db.Connect(gorm.Open)
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.GET("/", GetHealthCheck([]ServiceChecker{
		{
			Name:          "database",
			HealthChecker: db.HealthcheckDB,
		},
	}))

	req, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code, "Should be OK")
}
