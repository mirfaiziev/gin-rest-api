package server

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBResult struct {
	result string
}

func InitDatabase() {

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

	var result string
	db.Raw("select 1+1 as res").Scan(&result)

	fmt.Println(result)
}
