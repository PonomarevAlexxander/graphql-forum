package service

import (
	"context"

	"github.com/PonomarevAlexxander/graphql-forum/internal/dto"
	"github.com/google/uuid"
)

type UserService interface {
	Create(ctx context.Context, user dto.UserCreateInputDto) (dto.UserDto, error)
	Get(ctx context.Context, id uuid.UUID) (dto.UserDto, error)
}

type PostService interface {
	Count(ctx context.Context) (uint, error)
	Get(ctx context.Context, id uuid.UUID) (dto.PostDto, error)
	GetRange(ctx context.Context, params dto.PostGetDto) ([]dto.PostDto, error)
	Create(ctx context.Context, post dto.PostCreateDto) (dto.PostDto, error)
	Edit(ctx context.Context, post dto.PostUpdateDto) (dto.PostDto, error)
}

type CommentService interface {
	CountByParentId(ctx context.Context, parentId uuid.UUID) (uint, error)
	CountByPostId(ctx context.Context, postId uuid.UUID) (uint, error)
	Get(ctx context.Context, id uuid.UUID) (dto.CommentDto, error)
	Create(ctx context.Context, comment dto.CommentCreateInput) (dto.CommentDto, error)
	Edit(ctx context.Context, comment dto.CommentUpdateInput) (dto.CommentDto, error)
}
