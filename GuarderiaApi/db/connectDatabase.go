package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DBConn *gorm.DB

func Init() *gorm.DB {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:Jvsta123@tcp(127.0.0.1:3306)/guarderia?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		panic("failed to connect to database")
	}

	fmt.Println("Connection to database established")
	return db
}

func Close() {
	sqlDB, err := DBConn.DB()
	if err != nil {
		panic("cannot close database connection")
	}

	sqlDB.Close()
	fmt.Println("Connection to database closed")
}
