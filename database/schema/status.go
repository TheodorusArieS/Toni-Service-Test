package schema

import (
	"toni-service-test/query"
)

type Status struct {
	Base
	// Id int `gorm:"primary_key;column:id"`
	Name string `gorm:"not_null;unique;column:name"`
	Description string `gorm:"not_null;column:description"`
}

func (Status) TableName() string{
	return "statuses"
}
func (Status) Pk() string{
	return "id"
}

func (s Status) Ref() string {
	return s.TableName() + "(" + s.Pk() + ")"
}

func (s Status) AddForeignKey(){

}

func (s Status) InsertDefaults(){
	// Database.Exec(`
	// INSERT INTO statuses(id,name,description)
	// VALUES	(1,'Register','User Register'),
	// 		(2,'Completed','User Register Complete')
	// 		(3,'Pending','Transaction Status Pending')
	// 		(4,'Success','Transaction Status Success')
	// `)
	Database.Exec(query.InsertStatusDefault)
}