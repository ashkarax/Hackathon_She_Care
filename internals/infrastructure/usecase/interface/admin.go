package interfaceusecase

import (
	requestmodel "shecare/internals/models/reqModels"
	responsemodel "shecare/internals/models/resModels"
)

type IAdminUseCase interface {
	AdminLogin( *requestmodel.AdminLoginData) (*responsemodel.AdminLoginRes, error)
}
