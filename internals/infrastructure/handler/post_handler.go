package handler

import (
	"net/http"
	interfaceusecase "shecare/internals/infrastructure/usecase/interface"
	requestmodel "shecare/internals/models/reqModels"
	responsemodel "shecare/internals/models/resModels"

	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	PostUseCase interfaceusecase.IPostUseCase
}

func NewPostHandler(postUseCase interfaceusecase.IPostUseCase) *PostHandler {
	return &PostHandler{PostUseCase: postUseCase}
}

func (u *PostHandler) NewPost(c *gin.Context) {
	var postData requestmodel.PostData
	if err := c.BindJSON(&postData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	result, err := u.PostUseCase.NewPostUser(&postData)
	if err != nil {
		response := responsemodel.Responses(http.StatusBadRequest, "can't add new Post", result, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := responsemodel.Responses(http.StatusOK, "Post added succesfully", result, nil)
	c.JSON(http.StatusOK, response)

}
