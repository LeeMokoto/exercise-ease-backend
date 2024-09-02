package dynamo

import (
	"fmt"

	clientModel "main/pkg/client"
	inputModel "main/pkg/input"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

const tableName = "EE-Client-Table"

func GetClient(id inputModel.Input) ([]clientModel.Client, error) {
	print("getting clients")
	dynamoSession := createDynamoSession()
	input := &dynamodb.QueryInput{
		TableName: aws.String(tableName),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":UserId": {
				S: aws.String(id.UserId),
			},
			// "ClientId": {
			// 	S: aws.String(id.ClientId),
			// },
		},
		KeyConditionExpression: aws.String("UserId = :UserId"),
	}
	print("getting items from table")
	print(input)
	client := []clientModel.Client{}
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
