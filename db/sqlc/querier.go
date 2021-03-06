// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"context"
)

type Querier interface {
	CreateComment(ctx context.Context, arg CreateCommentParams) (Comment, error)
	CreatePost(ctx context.Context, arg CreatePostParams) (Post, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteComment(ctx context.Context, id int64) error
	DeletePost(ctx context.Context, id int64) error
	GetPost(ctx context.Context, id int64) (Post, error)
	GetUser(ctx context.Context, username string) (User, error)
	ListCommentsByPost(ctx context.Context, postID int64) ([]Comment, error)
	ListPosts(ctx context.Context, arg ListPostsParams) ([]Post, error)
	UpdateComment(ctx context.Context, arg UpdateCommentParams) (Comment, error)
	UpdatePost(ctx context.Context, arg UpdatePostParams) (Post, error)
}

var _ Querier = (*Queries)(nil)
