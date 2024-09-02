// main package
package main

// import other packages
import (
	"encoding/json"
	"net/http"

	clientModel "updateClient/main/pkg/client"
	"updateClient/main/pkg/db"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {

	// Start takes a handler and talks to an internal Lambda endpoint to pass requests to the handler.
	lambda.Start(handler)
}

func handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	postgresConnector := db.PostgresConnector{}
	db2, err := postgresConnector.GetConnection()
	if err != nil {
		return response(err.Error(), http.StatusBadRequest), nil
	}
	db2.AutoMigrate(&clientModel.Client{})
	var client clientModel.Client
	err = json.Unmarshal([]byte(req.Body), &client)

	if err != nil {
		return response(err.Error(), http.StatusBadRequest), nil
	}
	result := db2.Where("client_id = ?", client.ClientId).Save(&client)

	//err := json.Unmarshal([]byte(req.Body), &client)

	if result.Error != nil {
		return response(result.Error.Error(), http.StatusInternalServerError), nil
	}

	//dynamoErr := dynamo.SaveClient(client)

	// if dynamoErr != nil {
	// 	return response(dynamoErr.Error(), http.StatusInternalServerError), nil
	// }

	return response("client updated successfully", http.StatusOK), nil
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
