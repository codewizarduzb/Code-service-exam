package service

import (
	pbp "code-service-exam/post-service-exam/genproto/post-proto"
	l "code-service-exam/post-service-exam/pkg/logger"
	"code-service-exam/post-service-exam/storage"
	"context"

	grpcClient "code-service-exam/post-service-exam/service/grpc_client"

	"github.com/jmoiron/sqlx"
)

// Post Service ...
type PostService struct {
	storage storage.IStorage
	logger  l.Logger
	client  grpcClient.IServiceManager
}

// NewPostService ...
func NewPostService(db *sqlx.DB, log l.Logger, client grpcClient.IServiceManager) *PostService {
	return &PostService{
		storage: storage.NewStoragePg(db),
		logger:  log,
		client:  client,
	}
}

// Create a new post implementation
func (s *PostService) CreatePost(ctx context.Context, req *pbp.Post) (*pbp.Post, error) {
	post, err := s.storage.Post().CreatePost(req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return post, nil
}

// Get a post by ID implementation
func (s *PostService) GetPostById(ctx context.Context, req *pbp.GetPostByIdReq) (*pbp.GetPostByIdRes, error) {
	post, err := s.storage.Post().GetPostById(req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return post, nil
}

// Get all posts by user_id implementation
func (s *PostService) GetPostsByUserId(ctx context.Context, req *pbp.GetPostsByUserIdReq) (*pbp.GetPostsByUserIdRes, error) {
	posts, err := s.storage.Post().GetPostsByUserId(req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return posts, nil
}

// Update a post content by post_id implementation
func (s *PostService) UpdatePost(ctx context.Context, req *pbp.UpdatePostReq) (*pbp.UpdatePostRes, error) {
	succes, err := s.storage.Post().UpdatePost(req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return succes, nil
}

// Delete a post by post_id implementation
func (s *PostService) DeletePost(ctx context.Context, req *pbp.DeletePostReq) (*pbp.DeletePostRes, error) {
	succes, err := s.storage.Post().DeletePost(req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return succes, nil
}
