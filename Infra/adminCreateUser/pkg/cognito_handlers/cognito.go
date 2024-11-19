package cognito

import (
	"context"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	cognitoInput "github.com/lambda-go/pkg/input"
)

type CognitoActions struct {
	CognitoClient *cognitoidentityprovider.Client
}

func SaveClient(client cognitoInput.CognitoInput) (*cognitoidentityprovider.AdminCreateUserOutput, error) {
	fmt.Println("starting cognito session")
	cognitoSession := createCognitoSession()
	fmt.Println("declaring input")
	input := &cognitoidentityprovider.AdminCreateUserInput{
		UserPoolId:             aws.String(client.UserPoolId),
		Username:               aws.String(client.UserName),
		TemporaryPassword:      aws.String("demopassword"),
		DesiredDeliveryMediums: types.DeliveryMediumTypeEmail.Values(),
		UserAttributes: []types.AttributeType{{Name: aws.String("email"), Value: aws.String(client.UserEmail)}, {Name: aws.String("email_verified"), Value: aws.String("True")}, {Name: aws.String("phone_number_verified"), Value: aws.String("True")}, {Name: aws.String("profile"), Value: aws.String(client.Role)},
			{Name: aws.String("phone_number"), Value: aws.String(client.PhoneNumber)}},
	}
	// input1 := &cognitoidentityprovider.AdminDeleteUserInput{
	// 	UserPoolId:             aws.String(client.UserPoolId),
	// 	Username:               aws.String(client.UserName),
	// }
	fmt.Println("creating user")
	//cognitoSession.AdminDeleteUser(context.TODO(), input1)
	res, err := cognitoSession.AdminCreateUser(context.TODO(), input)

	if err != nil {
		var userExists *types.UsernameExistsException
		if errors.As(err, &userExists) {
			fmt.Println("User %v already exists in the user pool.")
		} else {
			fmt.Println("Couldn't create user %v. Here's why: %v\n")
		}
	}

	return res, nil
}

func createCognitoSession() *cognitoidentityprovider.Client {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("eu-west-1"),
	)
	if err != nil {
		fmt.Println("could not configure")
	}
	// sess := session.Must(session.NewSessionWithOptions(
	// 	session.Options{
	// 		SharedConfigState: session.SharedConfigEnable,
	// 	},
	// ))
	return cognitoidentityprovider.NewFromConfig(cfg)
}
