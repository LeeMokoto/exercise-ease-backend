// main package
package main

// import other packages
import (
	"encoding/json"
	clientModel "main/pkg/client"
	"main/pkg/db"
	inputModel "main/pkg/input"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler)
}

func handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	postgresConnector := db.PostgresConnector{}
	db2, err := postgresConnector.GetConnection()
	if err != nil {
		return errorResponse(err.Error(), http.StatusBadRequest), nil
	}
	var id inputModel.Input
	var client []clientModel.Client
	var reqId = req.QueryStringParameters["UserId"]
	var orgId = req.QueryStringParameters["OrgId"]
	//id.ClientId = clientId
	id.UserId = reqId + " - " + orgId
	print(id.UserId)
	// err := json.Unmarshal([]byte(req.Body), &id)

	// if err != nil {
	// 	return errorResponse("Couldn't unmarshal json into client struct", http.StatusBadRequest), nil
	// }

	//client, dynamoErr := dynamo.GetClient(id)
	result := db2.Find(&client, "user_id = ?", id.UserId)
	if result.Error != nil {
		return errorResponse(result.Error.Error(), http.StatusInternalServerError), nil
	}
	res, err := json.Marshal(client)
	if err != nil {
		return errorResponse(err.Error(), http.StatusInternalServerError), nil
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
