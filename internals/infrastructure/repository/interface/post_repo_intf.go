package interfacerepository

import requestmodel "shecare/internals/models/reqModels"

type IPostRepository interface {
	AddNewPost(*requestmodel.PostData) error
}
