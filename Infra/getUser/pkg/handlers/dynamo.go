package dynamo

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	userModel "main/pkg/user"
	inputModel "main/pkg/input"
)

const tableName = "EE-User-Table"

func GetUser(id inputModel.Input) (userModel.User,error) {
	dynamoSession := createDynamoSession()
	fmt.Println(id.UserId)
	fmt.Println(id.OrganisationID)
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
	user := userModel.User{}

	result, readErr := dynamoSession.GetItem(input)

	if readErr != nil {
		fmt.Println(readErr)
	} else {
			readErr = dynamodbattribute.UnmarshalMap(result.Item, &user)
			if readErr != nil {
    			panic(fmt.Sprintf("Failed to unmarshal Record, %v", readErr))
				}
		}
	return user, readErr
}

func createDynamoSession() *dynamodb.DynamoDB {
	sess := session.Must(session.NewSessionWithOptions(
		session.Options{
			SharedConfigState: session.SharedConfigEnable,
		},
	))

	return dynamodb.New(sess)
}
