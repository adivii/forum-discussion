package entities

import (
	forumEntities "adivii/forum-discussion/forum/entities"
	"time"
)

type (
	InsertPostDto struct {
		Id             int32               `gorm:"primaryKey;autoIncrement" json:"id"`
		ForumId        int32               `json:"forum_id"`           // Foreign Key
		Forum          forumEntities.Forum `gorm:"foreignKey:ForumId"` // Relation (Foreign at ForumId)
		Content        string              `json:"content"`
		IsOriginalPost bool                `json:"is_original_post"`
		IsReply        bool                `json:"is_reply"`
		RepliedId      int32               `json:"replied_id"` // Foreign Key
		Post           Post                `gorm:"foreignKey:RepliedId"`
		CreatedAt      time.Time           `json:"created_at"`
		UpdatedAt      time.Time           `json:"updated_at"`
		IsDeleted      bool                `json:"is_deleted"`
	}

	Post struct {
		Id             int32     `json:"id"`
		ForumId        int32     `json:"forum_id"`
		Content        string    `json:"content"`
		IsOriginalPost bool      `json:"is_original_post"`
		IsReply        bool      `json:"is_reply"`
		RepliedId      int32     `json:"replied_id"`
		CreatedAt      time.Time `json:"created_at"`
		UpdatedAt      time.Time `json:"updated_at"`
		IsDeleted      bool      `json:"is_deleted"`
	}

	PostResponseDto struct {
		Id             int32     `json:"id"`
		ForumId        int32     `json:"forum_id"`
		Content        string    `json:"content"`
		IsOriginalPost bool      `json:"is_original_post"`
		IsReply        bool      `json:"is_reply"`
		RepliedId      int32     `json:"replied_id"`
		CreatedAt      time.Time `json:"created_at"`
		UpdatedAt      time.Time `json:"updated_at"`
	}
)
