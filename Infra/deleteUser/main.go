// main package
package main

// import other packages
import (
	"deleteUser/main/pkg/db"
	inputModel "deleteUser/main/pkg/input"
	userModel "deleteUser/main/pkg/user"
	"fmt"
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
	var user userModel.User
	fmt.Println("this is getUser")
	var reqId = req.QueryStringParameters["UserId"]
	var orgId = req.QueryStringParameters["OrganisationID"]
	id.OrganisationID = orgId
	id.UserId = reqId
	//err := json.Unmarshal([]byte(req.Body), &id)
	fmt.Println(id)
	// if err != nil {
	// 	return errorResponse("Couldn't unmarshal json into user struct", http.StatusBadRequest), nil
	// }

	//user, dynamoErr := dynamo.GetUser(id)
	result := db2.Unscoped().Where("user_id = ? AND organisation_id = ?", id.UserId, id.OrganisationID).Delete(&user)
	if result.Error != nil {
		return errorResponse(result.Error.Error(), http.StatusInternalServerError), nil
	}

	return response("user deleted successfully", http.StatusOK), nil
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

func errorResponse(body string, statusCode int) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Body:       string(body),
		Headers: map[string]string{
			"Access-Control-Allow-Origin": "*",
		},
	}
}
