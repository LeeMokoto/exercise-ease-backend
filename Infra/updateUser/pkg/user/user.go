// This code is part of player package
package userModel

import (
	"time"

	"gorm.io/gorm"
)

// import other packages

type User struct {
	UserId         string `gorm:"uniqueIndex"`
	Name           string `json:"Name"`
	Surname        string `json:"Surname"`
	Email          string `json:"Email"`
	OrganisationID string `json:"OrganisationID"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt
}
