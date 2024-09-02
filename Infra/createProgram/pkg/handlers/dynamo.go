package dynamo

import (
	"fmt"

	programModel "createProgram/main/pkg/program"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

const programTableName = "EE-Programs-Table"

//const exerciseTableName = "EE-Exercise-Table"

func SaveProgam(program programModel.Program) error {
	programMap, marshalErr := dynamodbattribute.MarshalMap(program)
	//exerciseMap, marshalErr := dynamodbattribute.MarshalMap(program)

	if marshalErr != nil {
		fmt.Println("Failed to marshal to dynamo map")
		return marshalErr
	}

	dynamoSession := createDynamoSession()

	input := &dynamodb.PutItemInput{
		Item:      programMap,
		TableName: aws.String(programTableName),
	}

	_, writeErr := dynamoSession.PutItem(input)

	if writeErr != nil {
		fmt.Println("Failed to write to dynamo")
		return writeErr
	}

	return nil
}

func createDynamoSession() *dynamodb.DynamoDB {
	sess := session.Must(session.NewSessionWithOptions(
		session.Options{
			SharedConfigState: session.SharedConfigEnable,
		},
	))

	return dynamodb.New(sess)
}
