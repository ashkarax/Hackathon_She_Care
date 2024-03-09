package interfaceusecase

import (
	requestmodel "shecare/internals/models/reqModels"
	responsemodel "shecare/internals/models/resModels"
)

type IUserUseCase interface {
	UserSignup(*requestmodel.UserSignup) (*responsemodel.UserSignup, error)
	UserLogin(requestmodel.UserLogin) (responsemodel.UserSignup, error)
}
