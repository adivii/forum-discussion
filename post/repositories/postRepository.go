package repositories

import "adivii/forum-discussion/post/entities"

type PostRepository interface {
	InsertPostData(in *entities.InsertPostDto) error
}
