package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetModelos(c *gin.Context) {
	// Handler logic
	c.JSON(http.StatusOK, gin.H{"message": "GetModelos"})
}

func GetOperacoes(c *gin.Context) {
	// Handler logic
	c.JSON(http.StatusOK, gin.H{"message": "GetOperacoes"})
}

func CreateOperacao(c *gin.Context) {
	// Handler logic
	c.JSON(http.StatusCreated, gin.H{"message": "CreateOperacao"})
}

func GetColaboradores(c *gin.Context) {
	// Handler logic
	c.JSON(http.StatusOK, gin.H{"message": "GetColaboradores"})
}
