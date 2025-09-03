package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        string          `json:"id" gorm:"type:char(36);not null;primary_key;uniqueIndex"`
	FirstName string          `json:"first_name" gorm:"char(50); not null"`
	LastName  string          `json:"last_name" gorm:"char(50); not null"`
	Email     string          `json:"email" gorm:"char(50); not null; unique_index"`
	Phone     string          `json:"phone" gorm:"char(30); not null; unique_index"`
	CreatedAt *time.Time      `json:"-"`
	UpdatedAt *time.Time      `json:"-"`
	Delete    *gorm.DeletedAt `json:"-"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == "" {
		u.ID = uuid.New().String()
	}
	return
}
