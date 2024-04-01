package service

import (
	pbc "code-service-exam/comment-service-exam/genproto/comment-proto"
	l "code-service-exam/comment-service-exam/pkg/logger"
	"code-service-exam/comment-service-exam/storage"
	"context"

	grpcClient "code-service-exam/comment-service-exam/service/grpc_client"

	"github.com/jmoiron/sqlx"
)

// Comment Service ...
type CommentService struct {
	storage storage.IStorage
	logger  l.Logger
	client  grpcClient.IServiceManager
}

// NewCommentService ...
func NewCommentService(db *sqlx.DB, log l.Logger, client grpcClient.IServiceManager) *CommentService {
	return &CommentService{
		storage: storage.NewStoragePg(db),
		logger:  log,
		client:  client,
	}
}

// Create a new comment implementation
func (s *CommentService) CreateComment(ctx context.Context, req *pbc.Comment) (*pbc.Comment, error) {
	comment, err := s.storage.Comment().CreateComment(req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return comment, nil
}

// Get a comment by comment_id implementation
func (s *CommentService) GetComment(ctx context.Context, req *pbc.GetCommentRequest) (*pbc.GetCommentResponse, error) {
	post, err := s.storage.Comment().GetComment(req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return post, nil
}

// Update comment by comment_id implementation
func (s *CommentService) UpdateComment(ctx context.Context, req *pbc.UpdateCommentRequest) (*pbc.UpdateCommentResponse, error) {
	success, err := s.storage.Comment().UpdateComment(req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return success, nil
}

// Delete a comment by comment_id implementation
func (s *CommentService) DeleteComment(ctx context.Context, req *pbc.DeleteCommentRequest) (*pbc.DeleteCommentResponse, error) {
	succes, err := s.storage.Comment().DeleteComment(req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return succes, nil
}

// List of posts by post_id implementation
func (s *CommentService) ListComments(ctx context.Context, req *pbc.ListCommentsRequest) (*pbc.ListCommentsResponse, error) {
	comments, err := s.storage.Comment().ListComments(req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return comments, nil
}
