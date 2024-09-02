package dynamo

import (
	"fmt"

	inputModel "getPrograms/main/pkg/input"
	programModel "getPrograms/main/pkg/program"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

const tableName = "EE-Programs-Table"

func GetPrograms(id inputModel.Input) ([]programModel.Program, error) {
	dynamoSession := createDynamoSession()
	input := &dynamodb.QueryInput{
		TableName: aws.String(tableName),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":ProgramId": {
				S: aws.String(id.ProgramId),
			},
			// "ClientId": {
			// 	S: aws.String(id.ClientId),
			// },
		},
		KeyConditionExpression: aws.String("ProgramId = :ProgramId"),
	}
	print("getting items from table")
	print(input)
	client := []programModel.Program{}
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
