package controllers

import (
	"math/rand"
	"time-tracking/internal/models"

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

func AdminUserExists(db *gorm.DB) bool {
	var count int64
	db.Model(&models.User{}).Where("is_admin = ?", true).Count(&count)

	return count > 0
}

func CreateDefaultAdminUser(db *gorm.DB) (string, models.User) {
	password := RandomString(10)
	hash, _ := EncryptPassword(password)
	var admin = models.User{
		Name:         "admin",
		IsAdmin:      true,
		PasswordHash: hash,
		Email:        "test@test.com",
	}

	db.Save(&admin)
	return password, admin
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
