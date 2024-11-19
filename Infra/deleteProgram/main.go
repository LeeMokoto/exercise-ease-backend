// main package
package main

// import other packages
import (
	"deleteProgram/main/pkg/db"
	inputModel "deleteProgram/main/pkg/input"
	programModel "deleteProgram/main/pkg/program"
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

	var input inputModel.Input
	var program programModel.Program
	//var exercise exerciseModel.Exercise
	var programId = req.QueryStringParameters["ProgramId"]
	var id = req.QueryStringParameters["Id"]
	input.ProgramId = programId
	input.Id = id
	print(programId)
	print(id)
	result := db2.Unscoped().Where("program_id = ? AND id = ?", input.ProgramId, input.Id).Delete(&program)

	if result.Error != nil {
		return errorResponse(result.Error.Error(), http.StatusInternalServerError), nil
	}

	return response("program deleted successfully", http.StatusOK), nil
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
