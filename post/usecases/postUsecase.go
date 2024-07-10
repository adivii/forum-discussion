package usecases

import "adivii/forum-discussion/post/models"

type PostUsecase interface {
	InsertForumData(in *models.AddPostData) error
}
