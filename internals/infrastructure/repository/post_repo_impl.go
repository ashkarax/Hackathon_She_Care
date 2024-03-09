package repository

import (
	"fmt"
	interfacerepository "shecare/internals/infrastructure/repository/interface"
	requestmodel "shecare/internals/models/reqModels"

	"gorm.io/gorm"
)

type PostRepo struct {
	DB *gorm.DB
}

func NewPostRepository(DB *gorm.DB) interfacerepository.IPostRepository {
	return &PostRepo{DB: DB}
}

func (d *PostRepo) AddNewPost(postData *requestmodel.PostData) error {
	query := "INSERT INTO posts (title,content) VALUES (?)"
	values := []interface{}{
		postData.Title,
		postData.Content,
	}

	err := d.DB.Exec(query, values).Error
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
