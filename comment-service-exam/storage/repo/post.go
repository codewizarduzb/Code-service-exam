package repo

import (
	pbc "code-service-exam/comment-service-exam/genproto/comment-proto"
)

// CommentStorateI ...
type CommentStorageI interface {
	CreateComment(req *pbc.Comment) (*pbc.Comment, error)
	GetComment(req *pbc.GetCommentRequest) (*pbc.GetCommentResponse, error)
	UpdateComment(req *pbc.UpdateCommentRequest) (*pbc.UpdateCommentResponse, error)
	DeleteComment(req *pbc.DeleteCommentRequest) (*pbc.DeleteCommentResponse, error)
	ListComments(req *pbc.ListCommentsRequest) (*pbc.ListCommentsResponse, error)
}
