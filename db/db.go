package db

import (
	"fmt"
	"os"

	"github.com/luthfiswees/hotbox/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func CreateDB(){
	// Configurations
	dbUsername := os.Getenv("HOTBOX_DB_USERNAME")
	dbPassword := os.Getenv("HOTBOX_DB_PASSWORD")
	dbHost     := os.Getenv("HOTBOX_DB_HOST")
	dbPort     := os.Getenv("HOTBOX_DB_PORT")
	dbName     := os.Getenv("HOTBOX_DATABASE")

	// Connect to DB
	db, err := gorm.Open("mysql", dbUsername + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("Failed to connect to MySQL host " + dbHost + " in port " + dbPort + " when creating DB!")
		return
	}

	// Create DB if not exist
	_, err = db.Raw("CREATE DATABASE IF NOT EXISTS " + dbName + ";").Rows()
	if err != nil {
		fmt.Println("Create Database " + dbName + " failed!")
		return
	}

	db.Close()
	return
}

func MigrateDB(){
	// Configurations
	dbUsername := os.Getenv("HOTBOX_DB_USERNAME")
	dbPassword := os.Getenv("HOTBOX_DB_PASSWORD")
	dbHost     := os.Getenv("HOTBOX_DB_HOST")
	dbPort     := os.Getenv("HOTBOX_DB_PORT")
	dbName     := os.Getenv("HOTBOX_DATABASE")

	// Connect to DB
	db, err := gorm.Open("mysql", dbUsername + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("Failed to connect to MySQL host " + dbHost + " in port " + dbPort + " when migrating tables!")
		return
	}

	// Migrate Model
	db.AutoMigrate(&model.ReportEntry{})
	db.Close()

	return
}

func GetDatabaseInstance() (*gorm.DB, error){
	// Configurations
	dbUsername := os.Getenv("HOTBOX_DB_USERNAME")
	dbPassword := os.Getenv("HOTBOX_DB_PASSWORD")
	dbHost     := os.Getenv("HOTBOX_DB_HOST")
	dbPort     := os.Getenv("HOTBOX_DB_PORT")
	dbName     := os.Getenv("HOTBOX_DATABASE")

	// Connect to DB
	db, err := gorm.Open("mysql", dbUsername + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("Failed to connect to MySQL host " + dbHost + " in port " + dbPort + " when migrating tables!")
		return nil, err
	}

	// Return instance
	return db, nil
}