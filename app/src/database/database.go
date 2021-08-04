package database

import (
	"fmt"
	"log"
	video "rest-api-app/src/domain/video"

	"sync"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var connection *gorm.DB
var doOnce sync.Once

func initConnection() *gorm.DB {

	prosgret_conname := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable",
		viper.Get("POSTGRES_HOST"),
		viper.Get("POSTGRES_PORT"),
		viper.Get("POSTGRES_USER"),
		viper.Get("POSTGRES_DB"),
		viper.Get("POSTGRES_PASSWORD"),
	)

	db, err := gorm.Open(postgres.Open(prosgret_conname), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(&video.VideoEntity{})

	return db
}

func GetDB() *gorm.DB {
	doOnce.Do(func() {
		connection = initConnection()
	})

	return connection
}
