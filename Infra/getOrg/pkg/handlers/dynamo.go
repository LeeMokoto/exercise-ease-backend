package dynamo

import (
	"fmt"

	inputModel "main/pkg/input"
	orgModel "main/pkg/org"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

const tableName = "EE-Org-Table"

func GetOrg(id inputModel.Input) (orgModel.Org, error) {
	dynamoSession := createDynamoSession()
	input := &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"OrganisationOwnerID": {
				S: aws.String(id.OrganisationOwnerId),
			},
			"OrganisationID": {
				S: aws.String(id.OrganisationID),
			},
		},
	}
	println(id.OrganisationID)
	println(id.OrganisationOwnerId)

	org := orgModel.Org{}

	result, readErr := dynamoSession.GetItem(input)
	println(result)
	if readErr != nil {
		fmt.Println(readErr)
	} else {
		readErr = dynamodbattribute.UnmarshalMap(result.Item, &org)
		if readErr != nil {
			panic(fmt.Sprintf("Failed to unmarshal Record, %v", readErr))
		}
	}
	return org, readErr
}

func createDynamoSession() *dynamodb.DynamoDB {
	sess := session.Must(session.NewSessionWithOptions(
		session.Options{
			SharedConfigState: session.SharedConfigEnable,
		},
	))

	return dynamodb.New(sess)
}
