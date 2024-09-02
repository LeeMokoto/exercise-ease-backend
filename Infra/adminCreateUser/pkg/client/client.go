// This code is part of player package
package clientModel

// import other packages

type Client struct {
	UserPoolId  string `json:"UserPoolId"`
	UserName    string `json:"UserName"`
	UserEmail   string `json:"UserEmail"`
	PhoneNumber string `json:"PhoneNumber"`
	Role        string `json:"Role"`
}
