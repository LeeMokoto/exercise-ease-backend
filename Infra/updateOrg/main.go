// main package
package main

// import other packages
import (
	"encoding/json"
	"net/http"

	"updateOrg/main/pkg/db"
	orgModel "updateOrg/main/pkg/org"

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
	db2.AutoMigrate(&orgModel.Org{})
	var org orgModel.Org
	err = json.Unmarshal([]byte(req.Body), &org)
	if err != nil {
		return response(err.Error(), http.StatusBadRequest), nil
	}
	result := db2.Where("organisation_id = ?", org.OrganisationID).Save(&org)

	if result.Error != nil {
		return response(result.Error.Error(), http.StatusInternalServerError), nil
	}

	// fmt.Println("this is createOrg")
	// err := json.Unmarshal([]byte(req.Body), &org)

	// if err != nil {
	// 	return response("Couldn't unmarshal json into org struct", http.StatusBadRequest), nil
	// }

	// dynamoErr := dynamo.SaveOrg(org)

	// if dynamoErr != nil {
	// 	return response(dynamoErr.Error(), http.StatusInternalServerError), nil
	// }

	return response("successfully created organisation", http.StatusOK), nil
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
