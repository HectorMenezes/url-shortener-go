package url

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/HectorMenezes/url-shortener-go/db"
	"github.com/HectorMenezes/url-shortener-go/models"
	"github.com/HectorMenezes/url-shortener-go/utils"
	"github.com/HectorMenezes/url-shortener-go/cache"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func SetupTestRoutes() *gin.Engine {
	db.Connect(gorm.Open)
    cache.Start()

	db.Migrate()

	gin.SetMode(gin.ReleaseMode)
	routes := gin.Default()
	return routes
}

func mockUrlCreate() string {
	url := models.Url{Url: "https://chess.com"}

	db.GetDB().Create(&url)
	return url.ID
}

func mockUrlDelete(id string) {
	var url models.Url
	db.GetDB().Delete(&url, id)
}

func TestGetAllUrls(t *testing.T) {

	idMockedUrl := mockUrlCreate()
	defer mockUrlDelete(idMockedUrl)

	router := SetupTestRoutes()
	router.GET("/", GetUrls)
	req, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code, "Should be 200")

}

func TestGetOneUrl(t *testing.T) {

	idMockedUrl := mockUrlCreate()
	defer mockUrlDelete(idMockedUrl)

	router := SetupTestRoutes()
	router.GET("/:urlId", GetUrl)

	substest := []struct {
		name            string
		expectedCode    int
		id              string
		containedString string
	}{{
		name:            "happy-path",
		expectedCode:    http.StatusTemporaryRedirect,
		id:              idMockedUrl,
		containedString: `<a href="https://chess.com">Temporary Redirect</a>.`,
	},
		{
			name:            "random-id",
			expectedCode:    http.StatusNotFound,
			id:              "1234567",
			containedString: `{"message":"Url not found!"}`,
		},
	}
	for _, test := range substest {
		t.Run(test.name, func(t *testing.T) {
			req, _ := http.NewRequest("GET", fmt.Sprintf("/%s", test.id), nil)

			response := httptest.NewRecorder()
			router.ServeHTTP(response, req)

			assert.Equal(t, test.expectedCode, response.Code, fmt.Sprintf("Expected: %d got %d", test.expectedCode, response.Code))
			assert.Equal(t, true, strings.Contains(response.Body.String(), test.containedString),
				"Body: %s\nExpected: %s", response.Body.String(), test.containedString)
		})
	}

}

func TestCreateUrl(t *testing.T) {
	router := SetupTestRoutes()
	router.POST("/", CreateUrl)

	substest := []struct {
		name         string
		expectedCode int
		payload      string
		expectedHash string
	}{{
		name:         "happy-path",
		expectedCode: 201,
		payload:      `{"url": "https://chess.com"}`,
		expectedHash: "1534637",
	},
		{
			name:         "payload-witout-url",
			expectedCode: http.StatusBadRequest,
			payload:      `123`,
			expectedHash: "",
		},
		{
			name:         "payload-with-invalid-url",
			expectedCode: http.StatusBadRequest,
			payload:      `{"url": "chess"}`,
			expectedHash: "",
		},
	}

	for _, test := range substest {
		t.Run(test.name, func(t *testing.T) {
			req, _ := http.NewRequest("POST", "/", strings.NewReader(test.payload))

			response := httptest.NewRecorder()
			router.ServeHTTP(response, req)

			var url models.UlrResponsePost
			json.Unmarshal(response.Body.Bytes(), &url)

			assert.Equal(t, test.expectedCode, response.Code, fmt.Sprintf("Expected: %d got %d", test.expectedCode, response.Code))
			if response.Code == 201 {
				assert.Equal(t, url.Link, fmt.Sprintf("%s/%s",
					utils.GetEnv("SHORTENER_BASE_URL", ""),
					test.expectedHash,
				))
				mockUrlDelete(url.Link[len(url.Link)-7:])
			}

		})
	}
}
