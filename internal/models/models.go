package models

import (
	"time"
)

type BaseEntity struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type Modelo struct {
	BaseEntity
	Name string `json:"name"`
}

type Operacao struct {
	BaseEntity
	Name     string `json:"name"`
	Time     int    `json:"time"` // in seconds
	ModeloID int    `json:"modelo_id"`
	Modelo   Modelo `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"modelo"`
}

type Colaborador struct {
	BaseEntity
	Name string `json:"name"`
}

type Execucao struct {
	BaseEntity
	ModeloID      int         `json:"modelo_id"`
	Modelo        Modelo      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"modelo"`
	OperacaoID    int         `json:"operacao_id"`
	Operacao      Operacao    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"operacao"`
	ColaboradorID int         `json:"colaborador_id"`
	Colaborador   Colaborador `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"colaborador"`
	PercProdutivo float32     `gorm:"type:numeric(5,2)" json:"perc_produtivo"`
}

type CreateExecucaoRequest struct {
	ModeloID      int `json:"modelo_id"`
	OperacaoID    int `json:"operacao_id"`
	ColaboradorID int `json:"colaborador_id"`
	Minutes       int `json:"minutes"`
	Amount        int `json:"amount"`
}
