// main package
package main

// import other packages
import (
	"encoding/json"
	"net/http"
"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	dynamo "github.com/lambda-go/pkg/handlers"
	userModel "github.com/lambda-go/pkg/user"
)

func main() {

	// Start takes a handler and talks to an internal Lambda endpoint to pass requests to the handler.
	lambda.Start(handler)
}

func handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var user userModel.User
	fmt.Println("this is createUser")
	err := json.Unmarshal([]byte(req.Body), &user)

	if err != nil {
		return response("Couldn't unmarshal json into user struct", http.StatusBadRequest), nil
	}

	dynamoErr := dynamo.SaveUser(user)

	if dynamoErr != nil {
		return response(dynamoErr.Error(), http.StatusInternalServerError), nil
	}
	
	return response(user.Name, http.StatusOK), nil
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
