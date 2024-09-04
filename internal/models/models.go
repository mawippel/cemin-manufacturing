package models

import "gorm.io/gorm"

type Modelo struct {
	gorm.Model
	Name string
}

type Operacao struct {
	gorm.Model
	Name     string
	Time     int // in seconds
	ModeloID int
	Modelo   Modelo `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Colaborador struct {
	gorm.Model
	Name string
}

type Execucao struct {
	gorm.Model
	ModeloID      int
	Modelo        Modelo `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	OperacaoID    int
	Operacao      Operacao `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ColaboradorID int
	Colaborador   Colaborador `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	PercProdutivo float32     `gorm:"type:numeric(5,2)"`
}
