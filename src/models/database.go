package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "root:Jv410551@tcp(localhost)/UrWay?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
			NoLowerCase:   true,
		},
	})

	if err != nil {
		panic("Failed to connect to database!")
	}

	err = db.AutoMigrate(&Usuario{})
	if err != nil {
		panic("Failed to migrate database!")
	}

	DB = db
}
