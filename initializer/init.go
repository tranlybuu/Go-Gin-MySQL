package initializer

import (
	"context"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	dbUrl := os.Getenv("MYSQL_DB_CONN_STR")
	DB, err = gorm.Open(mysql.Open(dbUrl), &gorm.Config{})
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
	return DB
}

func ConnectMongoDB() *mongo.Client {
	ConnectEnv()
	uri := os.Getenv("MONGO_DB_CONN_STR")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	return client
}
