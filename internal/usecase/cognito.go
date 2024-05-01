package usecase

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	_ "github.com/joho/godotenv/autoload"
	"os"
)

var (
	AwsRegion            = os.Getenv("AWS_REGION")
	AwsCognitoClientID   = os.Getenv("AWS_COGNITO_CLIENT_ID")
	AwsCognitoUserPoolID = os.Getenv("AWS_COGNITO_USER_POOL_ID")
)

type CognitoClient interface {
	SignUp(email, username, password string) (error, string)
	ConfirmSignUp(email string, code string) (error, string)
	SignIn(email string, password string) (error, string, *cognito.InitiateAuthOutput)
	RefreshToken(refreshToken string) (error, string, *cognito.InitiateAuthOutput)
}

type awsCognitoClient struct {
	cognitoClient *cognito.CognitoIdentityProvider
	appClientId   string
}

func NewCognitoClient(cognitoRegion string, cognitoAppClientId string) CognitoClient {
	conf := &aws.Config{
		Region: aws.String(cognitoRegion),
	}

	sess, err := session.NewSession(conf)
	client := cognito.New(sess)
	if err != nil {
		panic(err)
	}

	return &awsCognitoClient{
		cognitoClient: client,
		appClientId:   cognitoAppClientId,
	}
}

func (ctx *awsCognitoClient) SignUp(email, username, password string) (error, string) {
	user := &cognito.SignUpInput{
		Username: aws.String(email),
		Password: aws.String(password),
		ClientId: aws.String(ctx.appClientId),
		UserAttributes: []*cognito.AttributeType{
			{
				Name:  aws.String("email"),
				Value: aws.String(email),
			},
			{
				Name:  aws.String("name"),
				Value: aws.String(username),
			},
		},
	}

	result, err := ctx.cognitoClient.SignUp(user)
	if err != nil {
		return err, ""
	}

	return nil, result.String()
}

func (ctx *awsCognitoClient) ConfirmSignUp(email string, code string) (error, string) {
	user := &cognito.ConfirmSignUpInput{
		Username:         aws.String(email),
		ClientId:         aws.String(ctx.appClientId),
		ConfirmationCode: aws.String(code),
	}

	result, err := ctx.cognitoClient.ConfirmSignUp(user)
	if err != nil {
		return err, ""
	}

	return nil, result.String()
}

func (ctx *awsCognitoClient) SignIn(email string, password string) (error, string, *cognito.InitiateAuthOutput) {
	user := &cognito.InitiateAuthInput{
		AuthFlow: aws.String("USER_PASSWORD_AUTH"),
		ClientId: aws.String(ctx.appClientId),
		AuthParameters: aws.StringMap(map[string]string{
			"USERNAME": email,
			"PASSWORD": password,
		}),
	}

	result, err := ctx.cognitoClient.InitiateAuth(user)
	if err != nil {
		return err, "", nil
	}

	return nil, result.String(), result
}

func (ctx *awsCognitoClient) RefreshToken(refreshToken string) (error, string, *cognito.InitiateAuthOutput) {
	user := &cognito.InitiateAuthInput{
		AuthFlow: aws.String("REFRESH_TOKEN_AUTH"),
		ClientId: aws.String(ctx.appClientId),
		AuthParameters: aws.StringMap(map[string]string{
			"REFRESH_TOKEN": refreshToken,
		}),
	}

	result, err := ctx.cognitoClient.InitiateAuth(user)
	if err != nil {
		return err, "", nil
	}

	return nil, result.String(), result
}
