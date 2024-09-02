// main package
package main

// import other packages
import (
	"encoding/json"
	"getPrograms/main/pkg/db"
	inputModel "getPrograms/main/pkg/input"
	programModel "getPrograms/main/pkg/program"
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
	var programs []programModel.Program
	//var exercise exerciseModel.Exercise
	var reqId = req.QueryStringParameters["UserId"]
	var orgId = req.QueryStringParameters["OrgId"]
	var clientId = req.QueryStringParameters["ClientId"]
	//id.ClientId = clientId
	id.ProgramId = reqId + " - " + orgId + " - " + clientId
	print(id.ProgramId)
	// err := json.Unmarshal([]byte(req.Body), &id)

	// if err != nil {
	// 	return errorResponse("Couldn't unmarshal json into client struct", http.StatusBadRequest), nil
	// }

	//client, dynamoErr := dynamo.GetPrograms(id)
	result := db2.Model(&programModel.Program{}).Preload("Exercises").Find(&programs, "program_id = ?", id.ProgramId)

	if result.Error != nil {
		return errorResponse(result.Error.Error(), http.StatusInternalServerError), nil
	}
	res, err := json.Marshal(programs)
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
