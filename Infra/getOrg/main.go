// main package
package main

// import other packages
import (
	"encoding/json"
	"fmt"
	"main/pkg/db"
	inputModel "main/pkg/input"
	orgModel "main/pkg/org"
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
	var org orgModel.Org
	fmt.Println("this is getOrg")
	var reqId = req.QueryStringParameters["OrganisationOwnerId"]
	var orgId = req.QueryStringParameters["OrganisationID"]
	id.OrganisationID = orgId
	id.OrganisationOwnerId = reqId
	//err := json.Unmarshal([]byte(req.Body), &id)

	// if err != nil {
	// 	return errorResponse("Couldn't unmarshal json into org struct", http.StatusBadRequest), nil
	// }

	//org, dynamoErr := dynamo.GetOrg(id)

	result := db2.Find(&org, "organisation_owner_id = ? AND organisation_id = ?", id.OrganisationOwnerId, id.OrganisationID)
	if result.Error != nil {
		return errorResponse(result.Error.Error(), http.StatusInternalServerError), nil
	}

	res, err := json.Marshal(org)
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
