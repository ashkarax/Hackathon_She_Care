package usecase

import (
	"shecare/internals/config"
	interfacerepository "shecare/internals/infrastructure/repository/interface"
	interfaceusecase "shecare/internals/infrastructure/usecase/interface"
	requestmodel "shecare/internals/models/reqModels"
	responsemodel "shecare/internals/models/resModels"
	helper "shecare/pkg"
)

type adminUsecase struct {
	repo             interfacerepository.IAdminRepository
	tokenSecurityKey *config.Config
}

func NewAdminUseCase(adminRepository interfacerepository.IAdminRepository, key config.Config) interfaceusecase.IAdminUseCase {
	return &adminUsecase{repo: adminRepository,
		tokenSecurityKey: &key}
}

func (r *adminUsecase) AdminLogin(adminData *requestmodel.AdminLoginData) (*responsemodel.AdminLoginRes, error) {
	var adminLoginRes responsemodel.AdminLoginRes

	HashedPassword, err := r.repo.GetPassword(adminData.Email)
	if err != nil {
		return nil, err
	}

	err = helper.CompairPassword(HashedPassword, adminData.Password)
	if err != nil {
		return nil, err
	}

	token, err := helper.GenerateRefreshToken(r.tokenSecurityKey.AdminSecurityKey)
	if err != nil {
		return nil, err
	}

	adminLoginRes.Token = token
	return &adminLoginRes, nil
}
