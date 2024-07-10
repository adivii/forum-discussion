package usecases

import (
	"adivii/forum-discussion/post/entities"
	"adivii/forum-discussion/post/models"
	"adivii/forum-discussion/post/repositories"
)

type PostUsecaseImpl struct {
	postRepository repositories.PostRepository
}

func NewPostUsecaseImpl(postRepository repositories.PostRepository) PostUsecase {
	return &PostUsecaseImpl{
		postRepository: postRepository,
	}
}

// InsertForumData implements PostUsecase.
func (p *PostUsecaseImpl) InsertForumData(in *models.AddPostData) error {
	insertPostData := &entities.InsertPostDto{
		ForumId:        in.ForumId,
		Content:        in.Content,
		IsOriginalPost: in.IsOriginalPost,
		IsReply:        in.IsReply,
		RepliedId:      in.RepliedId,
	}

	if err := p.postRepository.InsertPostData(insertPostData); err != nil {
		return err
	}

	return nil
}
