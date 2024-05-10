package schemas

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID       uint    `gorm:"primaryKey;index" json:"id"`
	Username string  `json:"username"`
	Password string  `json:"password"`
	Email    *string `json:"email"`
	Phone    *string `json:"phone"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
