package inmem

import (
	"context"
	"fmt"
	"slices"

	"github.com/PonomarevAlexxander/graphql-forum/internal/domain"
	"github.com/PonomarevAlexxander/graphql-forum/internal/errlib"
	"github.com/google/uuid"
)

type inmemCommentRepository struct {
	db *db
}

func NewInMemCommentRepository(db *db) *inmemCommentRepository {
	return &inmemCommentRepository{
		db: db,
	}
}

func (r *inmemCommentRepository) GetById(ctx context.Context, id uuid.UUID) (*domain.Comment, error) {
	comment, ok := r.db.comments[id]
	if !ok {
		return nil, errlib.ErrNotFound
	}
	return &comment, nil
}

func (r *inmemCommentRepository) GetByIds(ctx context.Context, ids []uuid.UUID) ([]domain.Comment, error) {
	comments := make([]domain.Comment, len(ids))
	for i, id := range ids {
		comment, ok := r.db.comments[id]
		if !ok {
			return nil, errlib.ErrNotFound
		}

		comments[i] = comment
	}

	return comments, nil
}

func (r *inmemCommentRepository) CountByParentId(ctx context.Context, parentId uuid.UUID) (uint, error) {
	return r.db.CountCommentsByCommentId(parentId)
}

func (r *inmemCommentRepository) CountByPostId(ctx context.Context, postId uuid.UUID) (uint, error) {
	return r.db.CountCommentsByPostId(postId)
}

func commentCompare(a, b domain.Comment) int {
	if a.Id.String() > b.Id.String() {
		return 1
	}
	if a.Id.String() < b.Id.String() {
		return -1
	}
	return 0
}

func commentWithIdCompare(c domain.Comment, u uuid.UUID) int {
	if c.Id.String() < u.String() {
		return -1
	}
	if c.Id.String() > u.String() {
		return 1
	}
	return 0
}

// GetRangeByParentId gets first limit results after provided id (after)
func (r *inmemCommentRepository) GetRangeByParentId(ctx context.Context, after *uuid.UUID, limit uint, parentId uuid.UUID) ([]domain.Comment, error) {
	comments, err := r.db.GetCommentComments(parentId)
	if err != nil {
		return nil, err
	}
	if len(comments) == 0 {
		return comments, nil
	}

	slices.SortFunc(comments, commentCompare)
	if after == nil {
		if limit > uint(len(comments)) {
			return comments, nil
		}
		return comments[:limit], nil
	}

	id, ok := slices.BinarySearchFunc(comments, *after, commentWithIdCompare)
	if !ok {
		return nil, fmt.Errorf("%w, comment not found", errlib.ErrNotFound)
	}
	if id == len(comments) {
		return []domain.Comment{}, nil
	}

	if limit+uint(id)+1 > uint(len(comments)) {
		return comments[id+1:], nil
	}
	return comments[id+1 : limit+uint(id)+1], nil
}

// GetRangeByPostId gets first limit results after provided id (after)
func (r *inmemCommentRepository) GetRangeByPostId(ctx context.Context, after *uuid.UUID, limit uint, postId uuid.UUID) ([]domain.Comment, error) {
	comments, err := r.db.GetPostComments(postId)
	if err != nil {
		return nil, err
	}

	slices.SortFunc(comments, commentCompare)
	if after == nil {
		if limit > uint(len(comments)) {
			return comments, nil
		}
		return comments[:limit], nil
	}

	id, ok := slices.BinarySearchFunc(comments, *after, commentWithIdCompare)
	if !ok {
		return nil, fmt.Errorf("%w, comment not found", errlib.ErrNotFound)
	}
	if id == len(comments) {
		return []domain.Comment{}, nil
	}

	if limit+uint(id)+1 > uint(len(comments)) {
		return comments[id+1:], nil
	}
	return comments[id+1 : limit+uint(id)+1], nil
}

func (r *inmemCommentRepository) GetRangeByParentIds(ctx context.Context, parentIds []uuid.UUID, limitPerParent uint) ([]domain.Comment, error) {
	comments := make([]domain.Comment, 0)
	for _, parentId := range parentIds {
		certainComments, err := r.GetRangeByParentId(ctx, nil, limitPerParent, parentId)
		if err != nil {
			return nil, err
		}

		for _, comment := range certainComments {
			comments = append(comments, comment)
		}
	}
	return comments, nil
}

func (r *inmemCommentRepository) GetRangeByPostIds(ctx context.Context, postIds []uuid.UUID, limitPerPost uint) ([]domain.Comment, error) {
	comments := make([]domain.Comment, 0)
	for _, postId := range postIds {
		certainComments, err := r.GetRangeByPostId(ctx, nil, limitPerPost, postId)
		if err != nil {
			return nil, err
		}

		for _, comment := range certainComments {
			comments = append(comments, comment)
		}
	}
	return comments, nil
}

func (r *inmemCommentRepository) CreateComment(ctx context.Context, comment domain.Comment) (*domain.Comment, error) {
	comment.Id = uuid.New()
	err := r.db.NewComment(comment)
	if err != nil {
		return nil, err
	}

	return &comment, nil
}

func (r *inmemCommentRepository) UpdateComment(ctx context.Context, comment domain.Comment) (*domain.Comment, error) {
	err := r.db.UpdateComment(comment)
	if err != nil {
		return nil, err
	}

	return &comment, nil
}
