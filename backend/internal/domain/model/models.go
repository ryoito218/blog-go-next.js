package model

import "time"

type User struct {
	ID           uint64 `gorm:"primaryKey"`
	Email        string `gorm:"size:255;not null;uniqueIndex"`
	PasswordHash string `gorm:"size:255;not null"`
	Role         string `gorm:"size:20;not null;default:user"`
	CreatedAt    time.Time
	UpdatedAt    time.Time

	Posts []Post `gorm:"foreignKey:UserID"`
}

type Post struct {
	ID          uint64     `gorm:"primaryKey"`
	UserID      uint64     `gorm:"not null;index"`
	Title       string     `gorm:"size:255;not null"`
	Slug        string     `gorm:"size:255;not null;uniqueIndex"`
	Content     string     `gorm:"type:longtext;not null"`
	Status      string     `gorm:"size:20;not null;default:draft"`
	PublishedAt *time.Time `gorm:"index"`
	CreatedAt   time.Time
	UpdatedAt   time.Time

	Images []PostImage `gorm:"foreignKey:PostID"`
}

type PostImage struct {
	ID        uint64 `gorm:"primaryKey"`
	PostID    uint64 `gorm:"not null;index"`
	Path      string `gorm:"size:1024;not null"`
	Alt       string `gorm:"size:255;not null;default:''"`
	CreatedAt time.Time
}