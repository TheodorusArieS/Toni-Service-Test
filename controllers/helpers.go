package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Helper struct {
	Context *gin.Context
}

func (h *Helper) OkResponse(data interface{}) {
	h.Context.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": data})
}

func (h *Helper) BadRequest(data interface{},message interface{}){
	h.Context.JSON(http.StatusBadRequest,gin.H{
		"status":http.StatusBadRequest,
		"data":data,
		"message":message,
	})
}

func (h *Helper) SuccessResponse(data interface{},message interface{}){
	h.Context.JSON(http.StatusOK,gin.H{
		"status":http.StatusOK,
		"data":data,
		"message:":message,
	})
}
