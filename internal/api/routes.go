package api

import (
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/modelos", GetModelos)
	router.POST("/modelos", CreateModelo)

	router.GET("/operacoes", GetOperacoes)
	router.POST("/modelos/:modeloId/operacoes", CreateOperacao)
	router.GET("/modelos/:modeloId/operacoes", GetOperacoesByModelo)

	router.POST("/colaboradores", CreateColaborador)
	router.GET("/colaboradores", GetColaboradores)

	router.POST("/execucoes", CreateExecucao)
	router.GET("/execucoes", GetExecucoes)

	return router
}
