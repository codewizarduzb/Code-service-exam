package service

import (
	pbu "code-service-exam/user-service-exam/genproto/user-proto"
	l "code-service-exam/user-service-exam/pkg/logger"
	"code-service-exam/user-service-exam/storage"
	"context"

	grpcClient "code-service-exam/user-service-exam/service/grpc_client"

	"github.com/jmoiron/sqlx"
)

// UserService ...
type UserService struct {
	storage storage.IStorage
	logger  l.Logger
	client  grpcClient.IServiceManager
}

// NewUserService ...
func NewUserService(db *sqlx.DB, log l.Logger, client grpcClient.IServiceManager) *UserService {
	return &UserService{
		storage: storage.NewStoragePg(db),
		logger:  log,
		client:  client,
	}
}

// Create User implementation
func (s *UserService) CreateUser(ctx context.Context, req *pbu.User) (*pbu.User, error) {
	user, err := s.storage.User().CreateUser(req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return user, nil
}

// GetUser implementation
func (s *UserService) GetUser(ctx context.Context, req *pbu.GetUserReq) (*pbu.User, error) {
	user, err := s.storage.User().GetUser(req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return user, nil
}

// ListUsers implementation
func (s *UserService) ListUsers(ctx context.Context, req *pbu.ListUsersRequest) (*pbu.ListUsersResponse, error) {
	users, err := s.storage.User().ListUsers(req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return users, nil
}

// UpdateUser implementation
func (s *UserService) UpdateUser(ctx context.Context, req *pbu.UpdateUserReq) (*pbu.UpdateUserRes, error) {
	user, err := s.storage.User().UpdateUser(req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return user, nil
}

// Delete User implementation
func (s *UserService) DeleteUser(ctx context.Context, req *pbu.DeleteUserReq) (*pbu.DeleteUserRes, error) {
	user, err := s.storage.User().DeleteUser(req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return user, nil
}
