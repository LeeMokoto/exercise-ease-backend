// main package
package main

// import other packages
import (
	"createProgram/main/pkg/db"
	exerciseModel "createProgram/main/pkg/exercise"
	programModel "createProgram/main/pkg/program"
	"encoding/json"
	"fmt"
	"net/http"

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
	db2.AutoMigrate(&programModel.Program{}, &exerciseModel.Exercise{})
	var program programModel.Program
	err = json.Unmarshal([]byte(req.Body), &program)
	fmt.Println(program)
	if err != nil {
		return response(err.Error(), http.StatusBadRequest), nil
	}
	result := db2.Create(&program)

	if result.Error != nil {
		return response(result.Error.Error(), http.StatusInternalServerError), nil
	}
	// var program programModel.Program
	// fmt.Println(req.Body)
	// err := json.Unmarshal([]byte(req.Body), &program)

	// if err != nil {
	// 	return response(err.Error(), http.StatusBadRequest), nil
	// }

	// dynamoErr := dynamo.SaveProgam(program)

	// if dynamoErr != nil {
	// 	return response(dynamoErr.Error(), http.StatusInternalServerError), nil
	// }

	return response("program added successfully", http.StatusOK), nil
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
