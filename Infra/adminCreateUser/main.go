// main package
package main

// import other packages
import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	cognito "github.com/lambda-go/pkg/cognito_handlers"
	cognitoInput "github.com/lambda-go/pkg/input"
)

type CognitoActions struct {
	CognitoClient *cognitoidentityprovider.Client
}

func main() {

	// Start takes a handler and talks to an internal Lambda endpoint to pass requests to the handler.
	lambda.Start(handler)
}

func handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var client cognitoInput.CognitoInput
	var user *cognitoidentityprovider.AdminCreateUserOutput
	fmt.Println("this is adminCreateClient")
	res := json.Unmarshal([]byte(req.Body), &client)
	if res != nil {
		return response(res.Error(), http.StatusBadRequest), nil
	}
	// var userPoolId = client.UserPoolId
	// var userName = client.UserName
	// var userEmail = client.UserEmail
	fmt.Println("initialising cognito call")
	user, err := cognito.SaveClient(client)
	if err != nil {
		return response(err.Error(), http.StatusBadRequest), nil
	}

	return response(*user.User.Username, http.StatusOK), nil
}

func response(body string, statusCode int) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Body:       string(body),
		Headers: map[string]string{
			"Access-Control-Allow-Origin": "*",
		},
	}
}
