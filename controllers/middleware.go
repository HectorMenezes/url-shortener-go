package url

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware verifies if the request has an Authorization header
func AuthMiddleware(c *gin.Context) {
	auth := c.Request.Header.Get("Authorization")
	if auth == "" {
		c.String(http.StatusForbidden, "No Authorization header provided")
		c.Abort()
		return
	}
}
