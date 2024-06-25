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
	orgModel "main/pkg/org"
	inputModel "main/pkg/input"
)

func main() {
	lambda.Start(handler)
}

func handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var id inputModel.Input
	var org orgModel.Org
	fmt.Println("this is getOrg")
	err := json.Unmarshal([]byte(req.Body), &id)

	if err != nil {
		return errorResponse("Couldn't unmarshal json into org struct", http.StatusBadRequest), nil
	}

	org, dynamoErr := dynamo.GetOrg(id)
	
	res, dynamoErr := json.Marshal(org)
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
