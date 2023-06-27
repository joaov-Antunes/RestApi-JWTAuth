package models

import "time"

type Tarefa struct {
	Id             uint      `json:"id" gorm:"primary_key"`
	Nome           string    `json:"nome"`
	DataVencimento time.Time `json:"DataVencimento"`
	DataCadastro   time.Time `json:"DataCadastro"`
	DataEdicao     time.Time `json:"DataEdicao"`
}
