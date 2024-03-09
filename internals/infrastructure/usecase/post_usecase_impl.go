package usecase

import (
	"fmt"
	interfacerepository "shecare/internals/infrastructure/repository/interface"
	interfaceusecase "shecare/internals/infrastructure/usecase/interface"
	requestmodel "shecare/internals/models/reqModels"
	responsemodel "shecare/internals/models/resModels"

	"github.com/go-playground/validator/v10"
)

type PostUseCase struct {
	PostRepo interfacerepository.IPostRepository
}

func NewPostUseCase(postRepo interfacerepository.IPostRepository) interfaceusecase.IPostUseCase {
	return &PostUseCase{PostRepo: postRepo}
}

func (r *PostUseCase) NewPostUser(postData *requestmodel.PostData) (*responsemodel.PostDataResp, error) {
	var respModel responsemodel.PostDataResp

	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(postData)
	if err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			for _, e := range ve {
				switch e.Field() {

				case "Title":
					respModel.Title = "validate:required,gte=4,lte=50"
				case "Content":
					respModel.Content = "validate:required,gte=30,lte=500"
				}
			}
			fmt.Println(err)
			return &respModel, err
		}
	}

	insertErr := r.PostRepo.AddNewPost(postData)
	if insertErr != nil {
		fmt.Println(insertErr)
		return &respModel, insertErr
	}
	return &respModel, nil

}
