package repository

import (
	"context"

	"github.com/PonomarevAlexxander/graphql-forum/internal/domain"
	"github.com/google/uuid"
)

type UserRepository interface {
	GetById(ctx context.Context, id uuid.UUID) (*domain.User, error)
	GetByIds(ctx context.Context, ids []uuid.UUID) ([]domain.User, error)
	GetByCommentIds(ctx context.Context, ids []uuid.UUID) ([]domain.UserWithCommentId, error)
	GetByPostIds(ctx context.Context, ids []uuid.UUID) ([]domain.UserWithPostId, error)
	CreateUser(ctx context.Context, user domain.User) (*domain.User, error)
}

type PostRepository interface {
	GetById(ctx context.Context, id uuid.UUID) (*domain.Post, error)
	GetByIds(ctx context.Context, ids []uuid.UUID) ([]domain.Post, error)
	GetRange(ctx context.Context, after *uuid.UUID, limit uint) ([]domain.Post, error)
	Count(ctx context.Context) (uint, error)
	CreatePost(ctx context.Context, post domain.Post) (*domain.Post, error)
	UpdatePost(ctx context.Context, post domain.Post) (*domain.Post, error)
}

type CommentRepository interface {
	GetById(ctx context.Context, id uuid.UUID) (*domain.Comment, error)
	GetByIds(ctx context.Context, ids []uuid.UUID) ([]domain.Comment, error)
	GetRangeByParentId(ctx context.Context, after *uuid.UUID, limit uint, parentId uuid.UUID) ([]domain.Comment, error)
	GetRangeByPostId(ctx context.Context, after *uuid.UUID, limit uint, postId uuid.UUID) ([]domain.Comment, error)
	GetRangeByParentIds(ctx context.Context, parentIds []uuid.UUID, limitPerParent uint) ([]domain.Comment, error)
	GetRangeByPostIds(ctx context.Context, postIds []uuid.UUID, limitPerPost uint) ([]domain.Comment, error)
	CountByParentId(ctx context.Context, parentId uuid.UUID) (uint, error)
	CountByPostId(ctx context.Context, postId uuid.UUID) (uint, error)
	CreateComment(ctx context.Context, post domain.Comment) (*domain.Comment, error)
	UpdateComment(ctx context.Context, comment domain.Comment) (*domain.Comment, error)
}
