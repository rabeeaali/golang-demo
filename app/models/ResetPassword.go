package models

import (
	"time"
)

type ResetPassword struct {
	ID        uint      `json:"id"`
	Email     string    `gorm:"type:varchar(191);not null" json:"email" form:"email"`
	Code      string    `gorm:"type:varchar(191);not null" json:"code" form:"code"`
	CreatedAt time.Time `json:"created_at"`
}
