package db

import (
	"fmt"
	"log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
	models "github.com/HectorMenezes/url-shortener-go/models"
	utils "github.com/HectorMenezes/url-shortener-go/utils"
)

var db *gorm.DB
var err error

// sqlOpener represents the type gorm.Open. 
type sqlOpener func(string, ...interface{}) (*gorm.DB, error)

// Connect performs the connectino to database and binds it
// to local var.
func Connect(open sqlOpener) error {
	log.Println("Starting Database...")

	db, err = open("postgres", fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		utils.GetEnv("POSTGRES_USER", ""),
		utils.GetEnv("POSTGRES_PASSWORD", ""),
		utils.GetEnv("POSTGRES_HOST", ""),
		utils.GetEnv("POSTGRES_PORT", ""),
		utils.GetEnv("POSTGRES_DB", ""),
	))

	if err != nil {
		log.Println("Failed to connect to database")
		return fmt.Errorf(err.Error())
	}

	log.Println("Database connected")
	return nil
}

// Migrate perferm all migrations based on models.
func Migrate() error {

	if db == (*gorm.DB)(nil) {
		return fmt.Errorf("database not connected")
	}

	if !db.HasTable(&models.Url{}) {
		db.CreateTable(&models.Url{})
	}
	db.AutoMigrate(&models.Url{})
	return nil
}

// GetDB return the current database connection.
func GetDB() *gorm.DB {
	return db
}

// HealthcheckDB is a routine to check health of service.
func HealthcheckDB() bool {
	return GetDB().DB().Ping() == nil
}
