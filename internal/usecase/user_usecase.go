package usecase

import (
	"github.com/SyamSolution/user-service/internal/model"
	"github.com/SyamSolution/user-service/internal/repository"
	cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"log"
)

type UserUsecase struct {
	UserRepo repository.UserPersister
}

type UserExecutor interface {
	CreateUser(user model.UserRequest) (int, error)
	ConfirmUser(user model.ConfirmCode) (error, string)
	LoginUser(user model.SignIn) (error, string, *cognito.InitiateAuthOutput)
	GetUserByEmail(email string) (model.UserResponse, error)
	RefreshToken(refreshToken string) (error, string, *cognito.InitiateAuthOutput)
}

func NewUserUsecase(userUsecase *UserUsecase) UserExecutor {
	return userUsecase
}

func (uc *UserUsecase) CreateUser(user model.UserRequest) (int, error) {
	userID, err := uc.UserRepo.CreateUser(user)
	if err != nil {
		return 0, err
	}

	cognitoClient := NewCognitoClient(AwsRegion, AwsCognitoClientID)
	err, _ = cognitoClient.SignUp(user.Email, user.Username, user.Password)
	if err != nil {
		return 0, err
	}

	return userID, nil
}

func (uc *UserUsecase) ConfirmUser(user model.ConfirmCode) (error, string) {
	cognitoClient := NewCognitoClient(AwsRegion, AwsCognitoClientID)
	err, result := cognitoClient.ConfirmSignUp(user.Email, user.Code)
	if err != nil {
		return err, ""
	}

	return nil, result
}

func (uc *UserUsecase) LoginUser(user model.SignIn) (error, string, *cognito.InitiateAuthOutput) {
	cognitoClient := NewCognitoClient(AwsRegion, AwsCognitoClientID)
	err, result, initiateAuthOutput := cognitoClient.SignIn(user.Email, user.Password)
	if err != nil {
		log.Println(err.Error())
		return err, "", nil
	}

	return nil, result, initiateAuthOutput
}

func (uc *UserUsecase) RefreshToken(refreshToken string) (error, string, *cognito.InitiateAuthOutput) {
	cognitoClient := NewCognitoClient(AwsRegion, AwsCognitoClientID)
	err, result, initiateAuthOutput := cognitoClient.RefreshToken(refreshToken)
	if err != nil {
		log.Println(err.Error())
		return err, "", nil
	}

	return nil, result, initiateAuthOutput
}

func (uc *UserUsecase) GetUserByEmail(email string) (model.UserResponse, error) {
	user, err := uc.UserRepo.GetUserByEmail(email)
	if err != nil {
		log.Println(err.Error())
		return model.UserResponse{}, err
	}

	response := model.UserResponse{
		UserID:      user.UserID,
		Username:    user.Username,
		Email:       user.Email,
		FullName:    user.FullName,
		PhoneNumber: user.PhoneNumber,
		Address:     user.Address,
		City:        user.City,
		Country:     user.Country,
		PostalCode:  user.PostalCode,
		NIK:         user.NIK,
	}

	return response, nil
}
