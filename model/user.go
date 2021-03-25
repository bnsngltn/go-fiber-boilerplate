package model

import (
	"errors"
	"time"

	"github.com/bnsngltn/go-fiber-boilerplate/config"
	"gorm.io/gorm"
)

type User struct {
	ID uint `gorm:"primaryKey"`

	Username string `gorm:"unique;not null;size:100"`
	Password string `gorm:"not null;size:100"`

	// Timestamps
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

// TableName defines the table for this struct during our database migrations
func (User) TableName() string {
	return config.SchemaGeneric + ".users"
}

// ValidateContents
func (user User) ValidateContents() error {
	if user.Username == "" {
		return errors.New("Username cannot be none")
	}
	if len(user.Password) < 6 {
		return errors.New("Password should be at least 6 characters")
	}
	return nil
}
