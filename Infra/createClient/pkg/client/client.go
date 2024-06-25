// This code is part of player package
package clientModel

// import other packages

type Client struct {
	ClientId         string `json:"ClientId"`
	Name           string `json:"Name"`
	Surname        string `json:"Surname"`
	Email          string `json:"Email"`
	UserId string `json:"UserId"`
	OrganisationId string `json:"OrganisationId"`
	Contact int `json:"Contact"`
	Notes string `json:"Notes"`
	Status string `json:"Status"`
	DateOfBirth string `json:"DateOfBirth"`
	DateJoined string `json:"DateJoined"`
	ProgramID string `json:"ProgramID"`
	ResourceBucket string `json:"ResourceBucket"`

}
