package grpcClient

import (
	"code-service-exam/comment-service-exam/config"
	pbp "code-service-exam/comment-service-exam/genproto/post-proto"
	"fmt"

	"google.golang.org/grpc"
)

type IServiceManager interface {
	PostService() pbp.PostServiceClient
}

type serviceManager struct {
	cfg         config.Config
	postService pbp.PostServiceClient
}

func New(cfg config.Config) (IServiceManager, error) {
	// dail to post-service
	connUser, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.UserServiceHost, cfg.UserServicePort),
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, fmt.Errorf("user service dail host: %s port : %d", cfg.UserServiceHost, cfg.UserServicePort)
	}

	return &serviceManager{
		cfg:         cfg,
		postService: pbp.NewPostServiceClient(connUser),
	}, nil
}

func (s *serviceManager) PostService() pbp.PostServiceClient {
	return s.postService
}
