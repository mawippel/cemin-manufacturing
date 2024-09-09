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
	var modelo models.Modelo
	if err := c.ShouldBindJSON(&modelo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.DB.Create(&modelo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, modelo)
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
	var operacao models.Operacao
	if err := c.ShouldBindJSON(&operacao); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var modelo models.Modelo
	if err := db.DB.First(&modelo, c.Param("modeloId")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Modelo not found"})
		return
	}

	operacao.Modelo = modelo

	if err := db.DB.Create(&operacao).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, operacao)
}

func GetColaboradores(c *gin.Context) {
	var colaborador []models.Colaborador
	if err := db.DB.Find(&colaborador).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, colaborador)
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

func GetExecucoes(c *gin.Context) {
	var execucoes []models.Execucao
	if err := db.DB.Find(&execucoes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, execucoes)
}

func CreateExecucao(c *gin.Context) {
	var execucaoRequest models.CreateExecucaoRequest
	if err := c.ShouldBindJSON(&execucaoRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var modelo models.Modelo
	if err := db.DB.First(&modelo, execucaoRequest.ModeloID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Modelo not found"})
		return
	}

	var operacao models.Operacao
	if err := db.DB.First(&operacao, execucaoRequest.OperacaoID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Operacao not found"})
		return
	}

	var colaborador models.Colaborador
	if err := db.DB.First(&colaborador, execucaoRequest.ColaboradorID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Colaborador not found"})
		return
	}

	requestedAmount := (execucaoRequest.Minutes * 60) / operacao.Time
	actualAmount := execucaoRequest.Amount
	percentualProdutivo := float32(actualAmount) / float32(requestedAmount) * 100

	execucao := models.Execucao{
		Modelo:        modelo,
		Operacao:      operacao,
		Colaborador:   colaborador,
		PercProdutivo: percentualProdutivo,
	}

	if err := db.DB.Create(&execucao).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, execucao)
}
