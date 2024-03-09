package handler

import (
	"net/http"
	interfaceusecase "shecare/internals/infrastructure/usecase/interface"
	requestmodel "shecare/internals/models/reqModels"
	helper "shecare/pkg"

	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	AdminUseCase interfaceusecase.IAdminUseCase
}

func NewAdminHandler(useCase interfaceusecase.IAdminUseCase) *AdminHandler {
	return &AdminHandler{AdminUseCase: useCase}
}

func (u *AdminHandler) AdminLogin(c *gin.Context) {
	var loginCredential requestmodel.AdminLoginData

	err := c.BindJSON(&loginCredential)
	if err != nil {
		c.JSON(http.StatusUnauthorized, err)
		return
	}

	data, err := helper.Validation(loginCredential)
	if err != nil {
		c.JSON(http.StatusBadRequest, data)
		return
	}

	result, err := u.AdminUseCase.AdminLogin(&loginCredential)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusOK, result)
	}
}
