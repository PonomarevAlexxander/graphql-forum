package impl

import (
	"context"
	"fmt"
	"time"

	"github.com/PonomarevAlexxander/graphql-forum/internal/converter"
	"github.com/PonomarevAlexxander/graphql-forum/internal/dto"
	"github.com/PonomarevAlexxander/graphql-forum/internal/errlib"
	"github.com/PonomarevAlexxander/graphql-forum/internal/repository"
	"github.com/google/uuid"
)

type commentService struct {
	cRepo repository.CommentRepository
	pRepo repository.PostRepository
}

func NewCommentService(repos *repository.Repositories) *commentService {
	return &commentService{
		cRepo: repos.Comment,
		pRepo: repos.Post,
	}
}

func (s *commentService) CountByParentId(ctx context.Context, parentId uuid.UUID) (uint, error) {
	return s.cRepo.CountByParentId(ctx, parentId)
}

func (s *commentService) CountByPostId(ctx context.Context, postId uuid.UUID) (uint, error) {
	return s.cRepo.CountByPostId(ctx, postId)
}

func (s *commentService) Get(ctx context.Context, id uuid.UUID) (dto.CommentDto, error) {
	comment, err := s.cRepo.GetById(ctx, id)
	if err != nil {
		return dto.CommentDto{}, err
	}

	return *converter.DomainToDtoComment(*comment), nil
}

func (s *commentService) Create(ctx context.Context, comment dto.CommentCreateInput) (dto.CommentDto, error) {
	if len(comment.Content) > 2000 {
		return dto.CommentDto{}, fmt.Errorf("%w, comment should have less than 2000 symbols", errlib.ErrBadRequest)
	}

	post, err := s.pRepo.GetById(ctx, comment.PostID)
	if err != nil {
		return dto.CommentDto{}, err
	}

	if !post.CommentsAllowed {
		return dto.CommentDto{}, errlib.ErrAccessDenied
	}

	newComment, err := s.cRepo.CreateComment(ctx, converter.DtoToDomainCommentInput(comment))
	if err != nil {
		return dto.CommentDto{}, err
	}

	return *converter.DomainToDtoComment(*newComment), nil
}

func (s *commentService) Edit(ctx context.Context, comment dto.CommentUpdateInput) (dto.CommentDto, error) {
	current, err := s.cRepo.GetById(ctx, comment.CommentID)
	if err != nil {
		return dto.CommentDto{}, err
	}

	if current.AuthorId != comment.UserID {
		return dto.CommentDto{}, errlib.ErrAccessDenied
	}

	if comment.Content != nil {
		current.Content = *comment.Content
	}

	current.EditedAt = time.Now().UTC()

	newComment, err := s.cRepo.UpdateComment(ctx, *current)
	if err != nil {
		return dto.CommentDto{}, err
	}

	return *converter.DomainToDtoComment(*newComment), nil
}
