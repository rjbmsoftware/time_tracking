package main

import (
	"math/rand"

	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/golang-jwt/jwt/v5"
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

func createDefaultAdminUser(db *gorm.DB) string {
	password := RandomString(10)
	salt := RandomString(10)
	hash, _ := EncryptPassword(password + salt)
	var admin = User{
		Name:         "admin",
		IsAdmin:      true,
		Salt:         salt,
		PasswordHash: hash,
	}

	db.Save(&admin)
	return password
}

func RandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

var secretKey = []byte("secret-key")

func CreateToken(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"email": email,
			"exp":   time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
