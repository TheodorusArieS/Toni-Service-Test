package schema

import (
	"time"
	"github.com/jinzhu/gorm"
)

var (
	Database *gorm.DB
)

type TableInterface interface {
	Pk() string
	Ref() string
	AddForeignKey()
	InsertDefaults()
}

type Base struct {
	Id int `gorm:"primary_key"`
	CreatedAt time.Time `gorm:"not_null;column:created_at;DEFAULT:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"not_null;column:updated_at;DEFAULT:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP`
	CreatedBy int `gorm:"DEFAULT:null"`
	UpdatedBy int `gorm:"DEFAULT:null"`
	DeletedAt *time.Time `sql:"index"` // ini kenapa mintanya sql : index ya
}

func AutoMigrate(database *gorm.DB){
	Database = database

	//drop table
	database.DropTableIfExists(
		// Status{},

	)
	// AUTO MIGRATE
	database.AutoMigrate(
		Status{},
		User{},
		OTPVerification{},
	)

	// insert default value
	Status{}.InsertDefaults()



}