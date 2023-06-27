package models

import "time"

type Usuario struct {
	Id           uint      `json:"id" gorm:"primary_key"`
	Nome         string    `json:"nome"`
	Email        string    `json:"email"`
	Senha        string    `json:"senha" gorm:"tableName:usuario"`
	DataCadastro time.Time `json:"DataCadastro"`
	DataEdicao   time.Time `json:"DataEdicao"`
}
