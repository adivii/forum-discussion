package entities

import "time"

type (
	InsertForumDto struct {
		Id        int32     `gorm:"primaryKey;autoIncrement" json:"id"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		IsDeleted bool      `json:"is_deleted"`
	}

	Forum struct {
		Id        int32     `json:"id"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		IsDeleted bool      `json:"is_deleted"`
	}

	ForumResponseDto struct {
		Id        int32     `json:"id"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)
