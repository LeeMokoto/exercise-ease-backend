package dynamo

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	clientModel "main/pkg/client"
	inputModel "main/pkg/input"
)

const tableName = "EE-Client-Table"

func GetClient(id inputModel.Input) (clientModel.Client,error) {
	dynamoSession := createDynamoSession()
	input := &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"UserId": {
				S: aws.String(id.UserId),
			},
			"OrganisationID": {
				S: aws.String(id.OrganisationID),
			},
		},
	}
	client := clientModel.Client{}

	result, readErr := dynamoSession.GetItem(input)

	if readErr != nil {
		fmt.Println(readErr)
	} else {
			readErr = dynamodbattribute.UnmarshalMap(result.Item, &client)
			if readErr != nil {
    			panic(fmt.Sprintf("Failed to unmarshal Record, %v", readErr))
				}
		}
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
