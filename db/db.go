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

func Connect() {
	log.Println("Starting Database...")

	user := utils.GetEnv("POSTGRES_USER", "")
	password := utils.GetEnv("POSTGRES_PASSWORD", "")
	host := utils.GetEnv("POSTGRES_HOST", "")
	port := utils.GetEnv("POSTGRES_PORT", "")
	database := utils.GetEnv("POSTGRES_DB", "")

	dbinfo := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		user,
		password,
		host,
		port,
		database,
	)
	fmt.Print(dbinfo)
	db, err = gorm.Open("postgres", dbinfo)
	if err != nil {
		log.Println("Failed to connect to database")
		panic(err)
	}

	log.Println("Database connected")
}

func Init() {
	Connect()
	if !db.HasTable(&models.Url{}) {
		err := db.CreateTable(&models.Url{})
		if err != nil {
			log.Println("Table already exists")
		}
	}
	db.AutoMigrate(&models.Url{})
}

func GetDB() *gorm.DB {
	return db
}
