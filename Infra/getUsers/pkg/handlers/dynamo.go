package dynamo

import (
	"fmt"

	inputModel "getUsers/main/pkg/input"
	userModel "getUsers/main/pkg/user"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

const tableName = "EE-User-Table"

func GetUsers(id inputModel.Input) ([]userModel.User, error) {
	print("getting users")
	dynamoSession := createDynamoSession()
	input := &dynamodb.QueryInput{
		TableName: aws.String(tableName),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":Email": {
				S: aws.String(id.Email),
			},
			":OrganisationID": {
				S: aws.String(id.OrganisationID),
			},
		},
		FilterExpression:       aws.String("Email <> :Email"),
		KeyConditionExpression: aws.String("OrganisationID = :OrganisationID"),
	}
	print("getting items from table")
	print(input)
	client := []userModel.User{}
	result, readErr := dynamoSession.Query(input)
	print("attempted to read from table")
	print(readErr)
	if readErr != nil {
		fmt.Println(readErr)
	} else {
		readErr = dynamodbattribute.UnmarshalListOfMaps(result.Items, &client)
		if readErr != nil {
			panic(fmt.Sprintf("Failed to unmarshal Record, %v", readErr))
		}
	}
	print("returning from getting item")
	return client, readErr
}

func createDynamoSession() *dynamodb.DynamoDB {
	sess := session.Must(session.NewSessionWithOptions(
		session.Options{
			SharedConfigState: session.SharedConfigEnable,
		},
	))

	return dynamodb.New(sess)
}
