package url

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ServiceChecker struct {
	Name          string
	HealthChecker func() bool
}

func GetHealthCheck(checkers []ServiceChecker) func(c *gin.Context) {
	m := make(map[string]bool)

	for _, function := range checkers {
		m[function.Name] = function.HealthChecker()
	}

	return func(c *gin.Context) {
		c.JSON(http.StatusOK, m)
	}

}
