// main package
package main

// import other packages
import (
	"encoding/json"
	"fmt"
	"getUsers/main/pkg/db"
	inputModel "getUsers/main/pkg/input"
	userModel "getUsers/main/pkg/user"
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
	var user []userModel.User
	fmt.Println("this is getUsers")
	var orgId = req.QueryStringParameters["OrganisationID"]
	var email = req.QueryStringParameters["Email"]
	id.OrganisationID = orgId
	id.Email = email
	//err := json.Unmarshal([]byte(req.Body), &id)
	fmt.Println(id)
	// if err != nil {
	// 	return errorResponse("Couldn't unmarshal json into user struct", http.StatusBadRequest), nil
	// }

	//user, dynamoErr := dynamo.GetUsers(id)

	// if dynamoErr != nil {
	// 	print(dynamoErr)
	// }

	result := db2.Find(&user, "email != ? AND organisation_id = ?", id.Email, id.OrganisationID)
	if result.Error != nil {
		return errorResponse(result.Error.Error(), http.StatusInternalServerError), nil
	}

	res, err := json.Marshal(user)
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
