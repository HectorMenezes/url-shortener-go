package url

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestMiddleware(t *testing.T) {

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	group := router.Group("/shortener", AuthMiddleware)

	group.GET("/", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	req, _ := http.NewRequest("GET", "/shortener/", nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, req)
	assert.Equal(t, http.StatusForbidden, response.Code, "Should be forbidden")
}
