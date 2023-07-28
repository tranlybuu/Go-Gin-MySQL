package initializer

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

func ConnectEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
}

var DB *gorm.DB

func ConnectDb() *gorm.DB {
	ConnectEnv()
	var err error
	dbUrl := os.Getenv("DB_CONN_STR")
	DB, err = gorm.Open(mysql.Open(dbUrl), &gorm.Config{})
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
	return DB
}
