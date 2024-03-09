package interfaceusecase

import (
	requestmodel "shecare/internals/models/reqModels"
	responsemodel "shecare/internals/models/resModels"
)

type IPostUseCase interface {
	NewPostUser(*requestmodel.PostData) (*responsemodel.PostDataResp, error)
}
