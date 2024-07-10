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
