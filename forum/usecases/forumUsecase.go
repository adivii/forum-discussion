package usecases

import "adivii/forum-discussion/forum/models"

type ForumUsecase interface {
	InsertForumData(in *models.AddForumData) error
}
