package api

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/modelos", GetModelos)
	router.GET("/operacoes", GetOperacoes)
	router.POST("/operacoes", CreateOperacao)
	router.GET("/colaboradores", GetColaboradores)
	return router
}
