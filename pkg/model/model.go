package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"github.com/lukaszozimek/organization_service/pkg/util"
)

var db *gorm.DB //database

func init() {

	e := godotenv.Load() //Load .env file
	if e != nil {
		fmt.Print(e)
	}

	username := util.GetEnvOrDefault("DB_USER", "'user'")

	password := util.GetEnvOrDefault("DB_PASS", "'password'")

	dbName := util.GetEnvOrDefault("DB_NAME", "'db_organization_service'")

	dbHost := util.GetEnvOrDefault("DB_HOST", "'127.0.0.1'")
	dbPort := util.GetEnvOrDefault("DB_PORT", "'5455'")
	dbUri := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbPort, username, dbName, password) //Build connection string
	fmt.Println(dbUri)

	conn, err := gorm.Open("postgres", dbUri)
	if err != nil {
		fmt.Print(err)
	}

	db = conn
	db.Debug().AutoMigrate(&Organization{}) //Database migration
}

//returns a handle to the DB object
func GetDB() *gorm.DB {
	return db
}
