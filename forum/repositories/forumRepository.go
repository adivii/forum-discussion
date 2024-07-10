package repositories

import "adivii/forum-discussion/forum/entities"

type ForumRepository interface {
	InsertForumData(in *entities.InsertForumDto) error
}
