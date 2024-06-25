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
	clientModel "github.com/lambda-go/pkg/client"
)

func main() {

	// Start takes a handler and talks to an internal Lambda endpoint to pass requests to the handler.
	lambda.Start(handler)
}

func handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var client clientModel.Client
	fmt.Println("this is createClient")
	err := json.Unmarshal([]byte(req.Body), &client)

	if err != nil {
		return response("Couldn't unmarshal json into client struct", http.StatusBadRequest), nil
	}

	dynamoErr := dynamo.SaveClient(client)

	if dynamoErr != nil {
		return response(dynamoErr.Error(), http.StatusInternalServerError), nil
	}

	return response("client added successfully", http.StatusOK), nil
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
