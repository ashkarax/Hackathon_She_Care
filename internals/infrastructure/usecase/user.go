package usecase

import (
	"errors"
	"fmt"
	"shecare/internals/config"
	interfacerepository "shecare/internals/infrastructure/repository/interface"
	interfaceusecase "shecare/internals/infrastructure/usecase/interface"
	requestmodel "shecare/internals/models/reqModels"
	responsemodel "shecare/internals/models/resModels"
	helper "shecare/pkg"
)

type userUseCase struct {
	token *config.Config
	repo  interfacerepository.IUserRepository
}

func NewUserUsecase(token *config.Config, repo interfacerepository.IUserRepository) interfaceusecase.IUserUseCase {
	return &userUseCase{token: token,
		repo: repo}
}

func (u *userUseCase) UserSignup(userData *requestmodel.UserSignup) (*responsemodel.UserSignup, error) {

	var resSignup *responsemodel.UserSignup
	fmt.Println("readh usecase", u.token)

	if isExist := u.repo.IsUserExist(userData.Email); isExist >= 1 {
		return nil, errors.New("user is exist try again , with another phone number")
	} else {
		userData.Password = helper.HashPassword(userData.Password)

		var err error
		resSignup, err = u.repo.CreateUser(userData)
		if err != nil {
			return nil, err
		}
	}

	accessToken, err := helper.GenerateAcessToken(u.token.UserSecurityKey, resSignup.ID)
	if err != nil {
		resSignup.Error = err.Error()
		return resSignup, err
	}

	refreshToken, err := helper.GenerateRefreshToken(u.token.UserSecurityKey)
	if err != nil {
		resSignup.Error = err.Error()
	}
	resSignup.AccessToken = accessToken
	resSignup.RefreshToken = refreshToken

	return resSignup, nil
}

func (u *userUseCase) UserLogin(loginCredential requestmodel.UserLogin) (responsemodel.UserSignup, error) {
	var resUserLogin responsemodel.UserSignup

	password, err := u.repo.FetchPasswordUsingEmail(loginCredential.Email)
	if err != nil {
		resUserLogin.Error = err.Error()
		return resUserLogin, err
	}

	err = helper.CompairPassword(password, loginCredential.Password)
	if err != nil {
		resUserLogin.Error = err.Error()
		return resUserLogin, err
	}

	userID, err := u.repo.FetchUserID(loginCredential.Email)
	if err != nil {
		resUserLogin.Error = err.Error()
		return resUserLogin, err
	}

	accessToken, err := helper.GenerateAcessToken(u.token.UserSecurityKey, userID)
	if err != nil {
		resUserLogin.Error = err.Error()
		return resUserLogin, err
	}

	refreshToken, err := helper.GenerateRefreshToken(u.token.UserSecurityKey)
	if err != nil {
		resUserLogin.Error = err.Error()
	}

	resUserLogin.AccessToken = accessToken
	resUserLogin.RefreshToken = refreshToken

	return resUserLogin, nil
}



