package repo

import (
	pbp "code-service-exam/post-service-exam/genproto/post-proto"
)

// UserStorageI ...
type PostStorageI interface {
	CreatePost(req *pbp.Post) (*pbp.Post, error)
	GetPostById(req *pbp.GetPostByIdReq) (*pbp.GetPostByIdRes, error)
	GetPostsByUserId(req *pbp.GetPostsByUserIdReq) (*pbp.GetPostsByUserIdRes, error)
	UpdatePost(req *pbp.UpdatePostReq) (*pbp.UpdatePostRes, error)
	DeletePost(req *pbp.DeletePostReq) (*pbp.DeletePostRes, error)
}
