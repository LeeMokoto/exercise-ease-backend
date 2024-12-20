// This code is part of player package
package clientModel

import (
	"time"

	"gorm.io/gorm"
)

// import other packages

type Client struct {
	ClientId       string `gorm:"primarykey,index"`
	Name           string `json:"Name"`
	Surname        string `json:"Surname"`
	Email          string `json:"Email"`
	UserId         string `json:"UserId"`
	Contact        string `json:"Contact"`
	Notes          string `json:"Notes"`
	Status         string `json:"Status"`
	DateOfBirth    string `json:"DateOfBirth"`
	DateJoined     string `json:"DateJoined"`
	ProgramID      string `json:"ProgramID"`
	ResourceBucket string `json:"ResourceBucket"`
	OrganisationID string `json:"OrganisationID"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt
}
