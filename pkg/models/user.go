package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

const cost = 10

// User
type User struct {
	Email     string    `json:"email" bson:"email"`
	CreatedAt time.Time `json:"-" bson:"createdAt"`
	Password  string    `json:"-" bson:"password"`
}

func (user *User) MakePassword(password string) {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), cost)
	user.Password = string(hash)
}

func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	return err == nil
}
