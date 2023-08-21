package url

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ServiceChecker struct {
	Name          string
	HealthChecker func() bool
}

// @Summary API's HealthChecker
// @BasePath /
// @Schemes
// @Description Return the services used by API and it's statuses
// @Tags HealthChecker
// @Accept json
// @Produce json
// @Success 200 {object} map[string]bool "Redirects to request URL"
// @Router /healthcheck [get]
func GetHealthCheck(checkers []ServiceChecker) func(c *gin.Context) {
	m := make(map[string]bool)

	for _, function := range checkers {
		m[function.Name] = function.HealthChecker()
	}

	return func(c *gin.Context) {
		c.JSON(http.StatusOK, m)
	}

}
