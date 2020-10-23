package schema

import (
	"toni-service-test/utilities"
	
)

type User struct {
	Base
	Phone string `gorm:"unique;not_null"`
	PinNumber string `gorm:"not_null"`
	StatusID int `gorm:"not_null"`
	Username string `gorm:"DEFAULT:null"`
	Password string `gorm:"DEFAULT:null"`
	Type string `gorm:"DEFAULT:null"`
	Role string `gorm:"DEFAULT:null"`
	Active int `gorm:"DEFAULT:0"`
	ShopID string `gorm:"DEFAULT:null"`
	Salt string `gorm:"DEFAULT:null"`
	Picture string `gorm:"DEFAULT:null"`
	LastLogin string `gorm:"DEFAULT:null"`
	Email string `gorm:"DEFAULT:null"`
}

func (u *User) HashPin() string{
	u.PinNumber = utilities.HashThePassword(u.PinNumber)
	return u.PinNumber
}

func (User) TableName() string {
	return "users"
}

func (User) Pk() string {
	return "id"
}

func (u User) Ref() string {
	return u.TableName() + "(" + u.Pk() +")"
}

func (u User) AddForeignKeys(){

}

func (u User) InsertDefaults(){
	
}