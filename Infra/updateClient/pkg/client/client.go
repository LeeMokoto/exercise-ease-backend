// This code is part of player package
package clientModel

import (
	"time"

	"gorm.io/gorm"
)

// import other packages

type Client struct {
	ClientId       string `gorm:"index,primarykey"`
	Name           string `json:"Name"`
	Surname        string `json:"Surname"`
	Email          string `json:"Email"`
	UserId         string `json:"UserId"`
	OrganisationID string `json:"OrganisationID"`
	Contact        string `json:"Contact"`
	Notes          string `json:"Notes"`
	Status         string `json:"Status"`
	DateOfBirth    string `json:"DateOfBirth"`
	DateJoined     string `json:"DateJoined"`
	ProgramID      string `json:"ProgramID"`
	ResourceBucket string `json:"ResourceBucket"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt
}
