package models

type Comment struct {
	CommentId string `json:"comment_id"`
	PostId    string `json:"post_id"`
	UserId    string `json:"user_id"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}

type CreateCommentRequest struct {
	CommentId string `json:"comment_id"`
	PostId    string `json:"post_id"`
	UserId    string `json:"user_id"`
	Content   string `json:"content"`
}

type CommentResponse struct {
	CommentId string `json:"comment_id"`
	PostId    string `json:"post_id"`
	UserId    string `json:"user_id"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
}

type UpdateCommentResponse struct {
	Success string `json:"success"`
}

type ListCommentsResponse struct {
	Comments []*CommentResponse `json:"comments"`
}
