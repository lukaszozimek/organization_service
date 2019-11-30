package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"os"
)

var db *gorm.DB //database

func init() {

	e := godotenv.Load() //Load .env file
	if e != nil {
		fmt.Print(e)
	}

	username := os.Getenv("db_user")
	if username == "" {
		username = "'user'"
	}
	password := os.Getenv("db_pass")
	if password == "" {
		password = "'password'"
	}
	dbName := os.Getenv("db_name")
	if dbName == "" {
		dbName = "'db_organization_service'"
	}
	dbHost := os.Getenv("db_host")
	if dbHost == "" {
		dbHost = "'127.0.0.1'"
	}

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password) //Build connection string
	fmt.Println(dbUri)

	conn, err := gorm.Open("postgres", dbUri)
	if err != nil {
		fmt.Print(err)
	}

	db = conn
	db.Debug().AutoMigrate(&model.Organization{}) //Database migration
}

//returns a handle to the DB object
func GetDB() *gorm.DB {
	return db
}
