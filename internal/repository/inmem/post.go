package inmem

import (
	"context"
	"fmt"
	"slices"

	"github.com/PonomarevAlexxander/graphql-forum/internal/domain"
	"github.com/PonomarevAlexxander/graphql-forum/internal/errlib"
	"github.com/google/uuid"
)

type inmemPostRepository struct {
	db *db
}

func NewInMemPostRepository(db *db) *inmemPostRepository {
	return &inmemPostRepository{
		db: db,
	}
}

func (r *inmemPostRepository) GetById(ctx context.Context, id uuid.UUID) (*domain.Post, error) {
	post, err := r.db.GetPost(id)
	if err != nil {
		return nil, err
	}

	return &post, nil
}

func (r *inmemPostRepository) GetByIds(ctx context.Context, ids []uuid.UUID) ([]domain.Post, error) {
	posts := make([]domain.Post, len(ids))
	var err error

	for i, id := range ids {
		posts[i], err = r.db.GetPost(id)

		if err != nil {
			return nil, err
		}
	}

	return posts, nil
}

func postCompare(a, b domain.Post) int {
	if a.Id.String() > b.Id.String() {
		return 1
	}
	if a.Id.String() < b.Id.String() {
		return -1
	}
	return 0
}

func postWithIdCompare(c domain.Post, u uuid.UUID) int {
	if c.Id.String() < u.String() {
		return -1
	}
	if c.Id.String() > u.String() {
		return 1
	}
	return 0
}

// GetRange gets first limit results after provided id (after)
//
// if after is nil, then returned first limit posts
func (r *inmemPostRepository) GetRange(ctx context.Context, after *uuid.UUID, limit uint) ([]domain.Post, error) {
	posts, err := r.db.GetPosts()
	if err != nil {
		return nil, err
	}
	if len(posts) == 0 {
		return posts, nil
	}

	slices.SortFunc(posts, postCompare)
	if after == nil {
		if limit > uint(len(posts)) {
			return posts, nil
		}
		return posts[:limit], nil
	}

	id, ok := slices.BinarySearchFunc(posts, *after, postWithIdCompare)
	if !ok {
		return nil, fmt.Errorf("%w, post not found", errlib.ErrNotFound)
	}
	if id == len(posts) {
		return []domain.Post{}, nil
	}

	if limit+uint(id)+1 > uint(len(posts)) {
		return posts[id+1:], nil
	}
	return posts[id+1 : limit+uint(id)+1], nil
}

func (r *inmemPostRepository) Count(ctx context.Context) (uint, error) {
	return r.db.CountPosts(), nil
}

func (r *inmemPostRepository) CreatePost(ctx context.Context, post domain.Post) (*domain.Post, error) {
	post.Id = uuid.New()

	err := r.db.NewPost(post)
	if err != nil {
		return nil, err
	}

	return &post, nil
}

func (r *inmemPostRepository) UpdatePost(ctx context.Context, post domain.Post) (*domain.Post, error) {
	err := r.db.UpdatePost(post)
	if err != nil {
		return nil, err
	}

	return &post, nil
}
