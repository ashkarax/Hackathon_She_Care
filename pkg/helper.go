package helper

import (
	"errors"
	"fmt"
	responsemodel "shecare/internals/models/resModels"

	"github.com/go-playground/validator"
	"golang.org/x/crypto/bcrypt"
)

func Validation(data interface{}) (*[]responsemodel.Errors, error) {
	var afterErrorCorection []responsemodel.Errors
	var result responsemodel.Errors
	validate := validator.New()

	err := validate.Struct(data)
	if err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			for _, e := range ve {

				err := fmt.Sprintf("%s should be  %s %s ", e.Field(), e.Tag(), e.Param())
				result = responsemodel.Errors{Err: err}

				afterErrorCorection = append(afterErrorCorection, result)
			}
		}
		return &afterErrorCorection, errors.New("doesn't fulfill the requirements")
	}
	return &afterErrorCorection, nil
}

func HashPassword(password string) string {

	HashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err, "problem at hashing ")
	}
	return string(HashedPassword)
}

func CompairPassword(hashedPassword string, plainPassword string) error {

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))

	if err != nil {
		return errors.New("passwords do not match")
	}

	return nil
}
