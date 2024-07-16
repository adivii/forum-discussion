package usecases

import (
	"adivii/forum-discussion/post/entities"
	"adivii/forum-discussion/post/models"
)

type PostUsecase interface {
	InsertForumData(in *models.AddPostData) error
	GetAllOriginalPost() ([]entities.PostResponseDto, error)
	GetPostById(id int) ([]entities.PostResponseDto, error)
}
