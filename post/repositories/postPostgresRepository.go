package repositories

import (
	"adivii/forum-discussion/database"
	"adivii/forum-discussion/post/entities"

	"github.com/labstack/gommon/log"
)

type PostPostgresRepository struct {
	db database.Database
}

func NewPostPostgresRepository(db database.Database) PostRepository {
	return &PostPostgresRepository{db: db}
}

// InsertPostData implements PostRepository.
func (p *PostPostgresRepository) InsertPostData(in *entities.InsertPostDto) error {
	data := &entities.Post{
		ForumId:        in.ForumId,
		Content:        in.Content,
		IsOriginalPost: in.IsOriginalPost,
		IsReply:        in.IsReply,
		RepliedId:      in.RepliedId,
	}

	result := p.db.GetDb().Create(data)

	if result.Error != nil {
		log.Errorf("InsertPostData: %v", result.Error)
		return result.Error
	}

	log.Debugf("InsertPostData: %v", result.RowsAffected)
	return nil
}

// GetAllOriginalPost implements PostRepository.
func (p *PostPostgresRepository) GetAllOriginalPost() ([]entities.PostResponseDto, error) {
	var posts []entities.Post
	var responses []entities.PostResponseDto

	result := p.db.GetDb().Where("is_original_post = ?", true).Find(&posts)
	if result.Error != nil {
		return nil, result.Error
	}

	for i := 0; i < int(result.RowsAffected); i++ {
		post := &posts[i]
		responses = append(responses, entities.PostResponseDto{
			Id:             post.Id,
			ForumId:        post.ForumId,
			Content:        post.Content,
			IsOriginalPost: post.IsOriginalPost,
			IsReply:        post.IsReply,
			RepliedId:      post.RepliedId,
			CreatedAt:      post.CreatedAt,
			UpdatedAt:      post.UpdatedAt,
		})
	}

	return responses, nil
}

// GetPostById implements PostRepository.
func (p *PostPostgresRepository) GetPostById(id int) ([]entities.PostResponseDto, error) {
	var post entities.Post
	var responses []entities.PostResponseDto

	result := p.db.GetDb().First(&post, id)
	if result.Error != nil {
		return nil, result.Error
	}

	responses = append(responses, entities.PostResponseDto{
		Id:             post.Id,
		ForumId:        post.ForumId,
		Content:        post.Content,
		IsOriginalPost: post.IsOriginalPost,
		IsReply:        post.IsReply,
		RepliedId:      post.RepliedId,
		CreatedAt:      post.CreatedAt,
		UpdatedAt:      post.UpdatedAt,
	})

	return responses, nil
}
