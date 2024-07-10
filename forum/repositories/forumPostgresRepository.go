package repositories

import (
	"adivii/forum-discussion/database"
	"adivii/forum-discussion/forum/entities"

	"github.com/labstack/gommon/log"
)

type ForumPostgresRepository struct {
	db database.Database
}

func NewForumPostgresRepository(db database.Database) ForumRepository {
	return &ForumPostgresRepository{db: db}
}

// InsertForumData implements ForumRepository.
func (f *ForumPostgresRepository) InsertForumData(in *entities.InsertForumDto) error {
	data := &entities.Forum{}

	result := f.db.GetDb().Create(data)

	if result.Error != nil {
		log.Errorf("InsertForumData: %v", result.Error)
		return result.Error
	}

	log.Debugf("InsertForumData: %v", result.RowsAffected)
	return nil
}
