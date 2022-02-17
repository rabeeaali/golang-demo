package models

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	ID        uint      `json:"id"`
	Name      string    `gorm:"type:varchar(191);not null" json:"name"`
	Email     string    `gorm:"type:varchar(191);unique;index;not null" json:"email"`
	Password  string    `gorm:"type:varchar(191);not null" json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"-"`
}

type ErrorAuth struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	Code     string `json:"code,omitempty"`
}

type LoginResult struct {
	UserInfo interface{} `json:"user_info"`
	Token    string      `json:"token"`
}

/*
=========================
 Helpers
=========================
*/

// HashPassword encrypts userControllers password
func (user *User) HashPassword(password []byte) {
	hashedPassword, _ := bcrypt.GenerateFromPassword(password, 8)
	user.Password = string(hashedPassword)
}

// CheckPassword check user's password
func (user *User) CheckPassword(providedPassword []byte) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), providedPassword)
}
