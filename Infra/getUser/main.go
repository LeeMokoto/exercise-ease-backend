// main package
package main

// import other packages
import (
	"encoding/json"
	"net/http"
		"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	dynamo "main/pkg/handlers"
	userModel "main/pkg/user"
	inputModel "main/pkg/input"
)

func main() {
	lambda.Start(handler)
}

func handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var id inputModel.Input
	var user userModel.User
	err := json.Unmarshal([]byte(req.Body), &id)

	if err != nil {
		return errorResponse("Couldn't unmarshal json into user struct", http.StatusBadRequest), nil
	}

	user, dynamoErr := dynamo.GetUser(id)
	
	res, dynamoErr := json.Marshal(user)
	if dynamoErr != nil {
		return errorResponse(dynamoErr.Error(), http.StatusInternalServerError), nil
	}
	fmt.Println(res)
	fmt.Println(string(res))
	return response(res, http.StatusOK), nil
}

func response(body []byte, statusCode int) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Body:       string(body),
		Headers: map[string]string{
			"Access-Control-Allow-Origin": "*",
		},
	}
}

func errorResponse(body string, statusCode int) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Body:       string(body),
		Headers: map[string]string{
			"Access-Control-Allow-Origin": "*",
		},
	}
}
