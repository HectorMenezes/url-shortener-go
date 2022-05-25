package db

import (
	"fmt"
	"testing"

	"github.com/HectorMenezes/url-shortener-go/models"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func TestGetDB(t *testing.T) {
	assert.Equal(t, (*gorm.DB)(nil), GetDB(), "Should be nil at first")
	Connect(gorm.Open)
	assert.NotEqual(t, (*gorm.DB)(nil), GetDB(), "Shouldn't be nil anymore")
}

func TestMigrate(t *testing.T) {
	db = (*gorm.DB)(nil)
	assert.Equal(t, fmt.Errorf("database not connected"), Migrate(), "Shouldn't be able to migrate due to connection")

	Connect(gorm.Open)

	GetDB().DropTable(&models.Url{})
	assert.Equal(t, nil, Migrate(), "Return null due to happy ending")

}

func TestConnect(t *testing.T) {

	substest := []struct {
		name           string
		openerFunction sqlOpener
		expectedError  error
	}{{
		name:           "Gorm-opener",
		openerFunction: gorm.Open,
		expectedError:  nil,
	},
		{
			name: "dummy-opener",
			openerFunction: func(string, ...interface{}) (db *gorm.DB, err error) {
				return nil, fmt.Errorf("Foo")
			},
			expectedError: fmt.Errorf("Foo"),
		},
	}

	for _, test := range substest {
		t.Run(test.name, func(t *testing.T) {
			err := Connect(test.openerFunction)
			assert.Equal(t, test.expectedError, err, fmt.Sprintf("Expect: %s, got %s", err, test.expectedError))
		})
	}

}
