package dynamo

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	orgModel "github.com/lambda-go/pkg/org"
)

const tableName = "EE-Org-Table"

func SaveOrg(org orgModel.Org) error {
	fmt.Println(org)
	orgMap, marshalErr := dynamodbattribute.MarshalMap(org)
	fmt.Println(orgMap)
	if marshalErr != nil {
		fmt.Println("Failed to marshal to dynamo map")
		return marshalErr
	}

	dynamoSession := createDynamoSession()

	input := &dynamodb.PutItemInput{
		Item:      orgMap,
		TableName: aws.String(tableName),
	}

	_, writeErr := dynamoSession.PutItem(input)

	if writeErr != nil {
		fmt.Println(writeErr)
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
