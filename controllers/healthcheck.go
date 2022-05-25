package url

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ServiceChecker represents a routine to check
// the health of a service.
type ServiceChecker struct {
	Name          string
	HealthChecker func() bool
}

// GetHealthCheck evaluates the state of service.
// Returns a gin's HandlerFunc with structure `{"service": "status"}`.
func GetHealthCheck(checkers []ServiceChecker) func(c *gin.Context) {
	m := make(map[string]bool)

	for _, function := range checkers {
		m[function.Name] = function.HealthChecker()
	}

	return func(c *gin.Context) {
		c.JSON(http.StatusOK, m)
	}

}
