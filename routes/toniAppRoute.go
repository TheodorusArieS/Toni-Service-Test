package routes

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"toni-service-test/controllers"
)

type ToniApp struct{
}

var (
	router = gin.Default()
)

func (t *ToniApp) Run(){
	t.routing()
}

func (t *ToniApp) routing(){
	defer router.Run()
	router.GET("/health", func (c *gin.Context){
		c.JSON(http.StatusOK,gin.H{
			"status":200,
			"Message":"All Good",
		})
	})
	toniController := controllers.ToniAppController{}

	toniAPI := router.Group("/toni")
	{
		toniAPI.POST("/register",toniController.Register)
		toniAPI.POST("/createOTP",toniController.CreateOTP)
		toniAPI.POST("/login",toniController.Login)
	}
}