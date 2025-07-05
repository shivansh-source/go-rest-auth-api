package models

import (
"gorm.io/gorm"
"time"
)

type Post struct {
    gorm.Model
    Title   string `json:"title"`
    Content string `json:"content"`
    UserID  uint   `json:"user_id"` // Foreign key: belongs to User
    CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
