// main package
package main

// import other packages
import (
	"encoding/json"
	"fmt"
	"net/http"

	"updateUser/main/pkg/db"
	userModel "updateUser/main/pkg/user"

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
	db2.AutoMigrate(&userModel.User{})
	var user userModel.User
	err = json.Unmarshal([]byte(req.Body), &user)
	fmt.Println(user)
	if err != nil {
		return response(err.Error(), http.StatusBadRequest), nil
	}
	result := db2.Where("user_id = ?", user.UserId).Save(&user)
	//dynamoErr := dynamo.SaveUser(user)

	// if dynamoErr != nil {
	// 	return response(dynamoErr.Error(), http.StatusInternalServerError), nil
	// }
	if result.Error != nil {
		return response(result.Error.Error(), http.StatusInternalServerError), nil
	}

	return response("user updated successfully", http.StatusOK), nil
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
