package handler

import (
	"net/http"
	"shecare/internals/config"
	interfaceusecase "shecare/internals/infrastructure/usecase/interface"
	requestmodel "shecare/internals/models/reqModels"
	helper "shecare/pkg"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUseCase interfaceusecase.IUserUseCase
	config      *config.Config
}

func NewUserHandler(userUseCase interfaceusecase.IUserUseCase, config config.Config) *UserHandler {
	return &UserHandler{userUseCase: userUseCase, config: &config}
}

func (u *UserHandler) UserSignup(c *gin.Context) {

	var userSignupData requestmodel.UserSignup

	if err := c.BindJSON(&userSignupData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data, err := helper.Validation(userSignupData)
	if err != nil {
		c.JSON(http.StatusBadRequest, data)
		return
	}

	resSignup, err := u.userUseCase.UserSignup(&userSignupData)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusOK, resSignup)
	}
}

func (u *UserHandler) UserLogin(c *gin.Context) {
	var loginCredential requestmodel.UserLogin
	if err := c.BindJSON(&loginCredential); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	data, err := helper.Validation(loginCredential)
	if err != nil {
		c.JSON(http.StatusBadRequest, data)
		return
	}

	result, err := u.userUseCase.UserLogin(loginCredential)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusOK, result)
	}
}
