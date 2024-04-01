package repo

import (
	pb "code-service-exam/user-service-exam/genproto/user-proto"
)

// UserStorageI ...
type UserStorageI interface {
	CreateUser(req *pb.User) (*pb.User, error)
	GetUser(req *pb.GetUserReq) (*pb.User, error)
	ListUsers(req *pb.ListUsersRequest) (*pb.ListUsersResponse, error)
	UpdateUser(req *pb.UpdateUserReq) (*pb.UpdateUserRes, error)
	DeleteUser(req *pb.DeleteUserReq) (*pb.DeleteUserRes, error)
}
