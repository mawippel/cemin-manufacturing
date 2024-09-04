package api

import (
	"controleConfeccao/internal/db"
	"controleConfeccao/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetModelos(c *gin.Context) {
	var modelos []models.Modelo
	err := db.DB.Find(&modelos).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, modelos)
}

func CreateModelo(c *gin.Context) {
	// Handler logic
	c.JSON(http.StatusCreated, gin.H{"message": "CreateModelo"})
}

func GetOperacoes(c *gin.Context) {
	var operacoes []models.Operacao
	err := db.DB.Find(&operacoes).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, operacoes)
}

func GetOperacoesByModelo(c *gin.Context) {
	modeloId := c.Param("modeloId")
	var operacoes []models.Operacao
	err := db.DB.Where("modelo_id = ?", modeloId).Find(&operacoes).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, operacoes)
}

func CreateOperacao(c *gin.Context) {
	// Handler logic
	c.JSON(http.StatusCreated, gin.H{"message": "CreateOperacao"})
}

func GetColaboradores(c *gin.Context) {
	var colaboradores []models.Colaborador
	err := db.DB.Find(&colaboradores).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, colaboradores)
}

func CreateColaborador(c *gin.Context) {
	var colaborador models.Colaborador
	if err := c.ShouldBindJSON(&colaborador); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.DB.Create(&colaborador).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, colaborador)
}
