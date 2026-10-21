package model

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Email     string    `json:"email" gorm:"uniqueIndex;not null;size:150"`
	Password  string    `json:"-" gorm:"not null"` // "-" = không trả về JSON
	FullName  string    `json:"full_name" gorm:"size:100"`
	Role      string    `json:"role" gorm:"default:customer"` // "admin" | "customer"
	CreatedAt time.Time `json:"created_at"`
}
