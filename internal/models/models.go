package models

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	Name string `json:"name"`
}

type User struct {
	gorm.Model
	Id           int    `gorm:"type:int;primary_key"`
	Name         string `gorm:"type:varchar(255)"`
	Email        string `gorm:"unique"`
	PasswordHash string `gorm:"type:varchar(255)"`
	IsAdmin      bool   `gorm:"type:bool"`
}
