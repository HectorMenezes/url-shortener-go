package url

import (
	"net/http"

	"github.com/HectorMenezes/url-shortener-go/db"

	"github.com/gin-gonic/gin"
)

func GetHealthCheck(c *gin.Context) {
	type test struct {
		attr int
	}
	var t test
	db.GetDB().Raw("SELECT 1").Scan(&t)

	if t.attr != 0 {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to access the database service.",
		})
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{
			"message": "Service OK!",
		})
	}

}
