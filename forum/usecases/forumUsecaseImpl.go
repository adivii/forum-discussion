package usecases

import (
	"adivii/forum-discussion/forum/entities"
	"adivii/forum-discussion/forum/models"
	"adivii/forum-discussion/forum/repositories"
)

type ForumUsecaseImpl struct {
	forumRepository repositories.ForumRepository
}

func NewForumUsecaseImpl(forumRepository repositories.ForumRepository) ForumUsecase {
	return &ForumUsecaseImpl{
		forumRepository: forumRepository,
	}
}

// InsertForumData implements ForumUsecase.
func (f *ForumUsecaseImpl) InsertForumData(in *models.AddForumData) error {
	insertForumData := &entities.InsertForumDto{}

	if err := f.forumRepository.InsertForumData(insertForumData); err != nil {
		return err
	}

	return nil
}
