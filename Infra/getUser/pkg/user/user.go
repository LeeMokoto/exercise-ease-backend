// This code is part of player package
package userModel

// import other packages

type User struct {
	UserId         string `json:"UserId"`
	Name           string `json:"Name"`
	Surname        string `json:"Surname"`
	Email          string `json:"Email"`
	OrganisationID string `json:"OrganisationID"`
}
