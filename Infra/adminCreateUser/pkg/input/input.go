package cognitoInput

type CognitoInput struct {
	UserPoolId  string `json:"UserPoolId"`
	UserName    string `json:"UserName"`
	UserEmail   string `json:"UserEmail"`
	PhoneNumber string `json:"PhoneNumber"`
	Role        string `json:"Role"`
}
