package database

import (
	"github.com/jinzhu/gorm"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"fmt"
	"log"
	"toni-service-test/database/schema"
)

var (
	Database *gorm.DB
	Client *sql.DB
)

func Connect(){
	username := os.Getenv("username")
	password :=os.Getenv("password")
	host := os.Getenv("host")
	dbSchema :=os.Getenv("dbSchema")
	dataSource := fmt.Sprintf("%s:%s@tcp(%s)/%s",username,password,host,dbSchema)

	db, err := gorm.Open("mysql",dataSource)

	if err != nil{
		log.Println("MySQL:",err)
	}

	Database = db
	Client = db.DB()

	schema.AutoMigrate(db)
	log.Println("Database successfully created")
}

