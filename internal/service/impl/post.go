package impl

import (
	"context"
	"time"

	"github.com/PonomarevAlexxander/graphql-forum/internal/converter"
	"github.com/PonomarevAlexxander/graphql-forum/internal/dto"
	"github.com/PonomarevAlexxander/graphql-forum/internal/errlib"
	"github.com/PonomarevAlexxander/graphql-forum/internal/repository"
	"github.com/google/uuid"
)

type postService struct {
	repo repository.PostRepository
}

func NewPostService(repo repository.PostRepository) *postService {
	return &postService{
		repo: repo,
	}
}

func (s *postService) Count(ctx context.Context) (uint, error) {
	return s.repo.Count(ctx)
}

func (s *postService) Get(ctx context.Context, id uuid.UUID) (dto.PostDto, error) {
	post, err := s.repo.GetById(ctx, id)
	if err != nil {
		return dto.PostDto{}, err
	}

	return converter.DomainToDtoPost(post), nil
}

func (s *postService) GetRange(ctx context.Context, params dto.PostGetDto) ([]dto.PostDto, error) {
	posts, err := s.repo.GetRange(ctx, params.After, params.Limit)
	if err != nil {
		return nil, err
	}

	return converter.DomainToDtoPosts(posts), nil
}

func (s *postService) Create(ctx context.Context, post dto.PostCreateDto) (dto.PostDto, error) {
	newPost, err := s.repo.CreatePost(ctx, converter.DtoToDomainInputPost(post))
	if err != nil {
		return dto.PostDto{}, err
	}

	return converter.DomainToDtoPost(newPost), nil
}

func (s *postService) Edit(ctx context.Context, post dto.PostUpdateDto) (dto.PostDto, error) {
	current, err := s.repo.GetById(ctx, post.PostId)
	if err != nil {
		return dto.PostDto{}, err
	}

	if current.AuthorId != post.UserId {
		return dto.PostDto{}, errlib.ErrAccessDenied
	}

	if post.CommentsAllowed != nil {
		current.CommentsAllowed = *post.CommentsAllowed
	}
	if post.Content != nil {
		current.Content = *post.Content
	}
	if post.Title != nil {
		current.Title = *post.Title
	}
	current.EditedAt = time.Now().UTC()

	newPost, err := s.repo.UpdatePost(ctx, *current)
	if err != nil {
		return dto.PostDto{}, err
	}

	return converter.DomainToDtoPost(newPost), nil
}
