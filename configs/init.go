package configs

import (

	"github.com/joho/godotenv"
)

func InitEnv(){
	_=godotenv.Load()

}