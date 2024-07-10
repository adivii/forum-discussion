package models

type AddPostData struct {
	ForumId        int32  `json:"forum_id"`
	Content        string `json:"content"`
	IsOriginalPost bool   `json:"is_original_post"`
	IsReply        bool   `json:"is_reply"`
	RepliedId      int32  `json:"replied_id"`
}
