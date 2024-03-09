package interfacerepository

import (
	requestmodel "shecare/internals/models/reqModels"
	responsemodel "shecare/internals/models/resModels"
)

type IUserRepository interface {
	IsUserExist(string) int
	CreateUser(*requestmodel.UserSignup) (*responsemodel.UserSignup, error)
	FetchPasswordUsingEmail( string) (string, error)
	FetchUserID( string) (string, error)
}
