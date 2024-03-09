package repository

import (
	"database/sql"
	"errors"
	"fmt"
	interfacerepository "shecare/internals/infrastructure/repository/interface"
	requestmodel "shecare/internals/models/reqModels"
	responsemodel "shecare/internals/models/resModels"

	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) interfacerepository.IUserRepository {
	return &userRepository{DB: db}
}

func (d *userRepository) IsUserExist(email string) int {
	var userCount int

	query := "SELECT COUNT(*) FROM users WHERE email=$1 AND status!=$2"
	err := d.DB.Raw(query, email, "delete").Row().Scan(&userCount)
	if err != nil {
		fmt.Println("Error for user exist, using same email in signup")
	}
	return userCount
}

func (d *userRepository) CreateUser(userDetails *requestmodel.UserSignup) (*responsemodel.UserSignup, error) {
	fmt.Println("--callded user adde.dc")

	var userData responsemodel.UserSignup
	query := "INSERT INTO users (name, email,  password) VALUES($1, $3, $4) RETURNING *"
	result := d.DB.Raw(query, userDetails.Name, userDetails.Email, userDetails.Password).Scan(&userData)
	if result.Error != nil {
		return nil, result.Error
	}
	return &userData, nil
}

func (d *userRepository) FetchPasswordUsingEmail(email string) (string, error) {
	var password string

	query := "SELECT password FROM users WHERE email=? AND status='active'"
	row := d.DB.Raw(query, email).Row()
	fmt.Println("--------", row)

	if row == nil {
		return "", errors.New("no user exist or you are blocked by admin")
	}

	err := row.Scan(&password)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("user does not exist or user get blocked")
		}
		return "", fmt.Errorf("error scanning row: %s", err)
	}
	return password, nil
}

func (d *userRepository) FetchUserID(Email string) (string, error) {
	var userID string

	query := "SELECT id FROM users WHERE Email=? AND status='active'"
	data := d.DB.Raw(query, Email).Row()

	if err := data.Scan(&userID); err != nil {
		return "", errors.New("fetching user id cause error")
	}
	return userID, nil
}
