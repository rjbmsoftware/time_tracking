package main

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func EncryptPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func ComparePassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func adminUserExists(db *gorm.DB) bool {
	var count int64
	db.Model(&User{}).Where("is_admin = ?", true).Count(&count)

	return count > 0
}

func createDefaultAdminUser(db *gorm.DB) {
	password := "someRandomPassword"
	salt := "salty"
	hash := salt + password
	var admin = User{
		Name:         "admin",
		IsAdmin:      true,
		Salt:         salt,
		PasswordHash: hash,
	}

	// log helpful message about username password and needed change
	db.Save(&admin)
}
