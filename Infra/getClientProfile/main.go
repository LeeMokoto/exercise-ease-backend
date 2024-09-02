// main package
package main

// import other packages
import (
	"encoding/json"
	"fmt"
	clientModel "getClientProfile/main/pkg/client"
	dynamo "getClientProfile/main/pkg/handlers"
	inputModel "getClientProfile/main/pkg/input"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler)
}

func handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var id inputModel.Input
	var client []clientModel.Client
	fmt.Println("this is getUser")
	var clientId = req.QueryStringParameters["ClientId"]
	id.ClientId = clientId
	//err := json.Unmarshal([]byte(req.Body), &id)
	fmt.Println(id)
	// if err != nil {
	// 	return errorResponse("Couldn't unmarshal json into user struct", http.StatusBadRequest), nil
	// }

	client, dynamoErr := dynamo.GetClient(id)

	res, dynamoErr := json.Marshal(client)
	if dynamoErr != nil {
		return errorResponse(dynamoErr.Error(), http.StatusInternalServerError), nil
	}

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
