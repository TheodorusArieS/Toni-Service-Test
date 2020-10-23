package app

import (
	// "fmt"
	// "github.com/gin-gonic/gin"
	"toni-service-test/routes"
	"toni-service-test/configs"
	"toni-service-test/database"
)



func Start(){
	configs.InitEnv()
	database.Connect()
	api := routes.ToniApp{}
	api.Run()
	
}