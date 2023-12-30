package service

import "github.com/szymon676/codehund/types"

type PostServicer interface {
	CreatePost(*types.Post) error
	GetPosts() ([]*types.Post, error)
	DeletePost() error
	GetPostsByUser(username string) ([]*types.Post, error)
}
