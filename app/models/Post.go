package models

import (
	"gorm.io/gorm"
	"time"
)

type Post struct {
	ID        uint      `json:"id"`
	Title     string    `gorm:"type:varchar(191);not null" json:"title"`
	Desc      string    `gorm:"type:varchar(191);not null" json:"description"`
	UserID    int       `gorm:"not null" json:"-"`
	User      User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"user"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"-"`
}

func (post *Post) Count(db *gorm.DB) int64 {
	var total int64
	db.Model(&Post{}).Count(&total)
	return total
}

func (post *Post) Take(db *gorm.DB, limit int, offset int) interface{} {
	var posts []Post
	db.Preload("User").Offset(offset).Limit(limit).Find(&posts)
	return posts
}
