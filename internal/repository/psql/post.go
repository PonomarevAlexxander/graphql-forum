package psql

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/PonomarevAlexxander/graphql-forum/internal/db"
	"github.com/PonomarevAlexxander/graphql-forum/internal/domain"
	"github.com/PonomarevAlexxander/graphql-forum/internal/errlib"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type psqlPostRepository struct {
	db *db.DbProvider
}

func NewPsqlPostRepository(db *db.DbProvider) *psqlPostRepository {
	return &psqlPostRepository{
		db: db,
	}
}

func (r *psqlPostRepository) GetById(ctx context.Context, id uuid.UUID) (*domain.Post, error) {
	var post domain.Post
	if err := r.db.GetContext(ctx, &post, `select * from post where id=$1`, id); err != nil {
		slog.Error(
			"Some error while trying to get Post",
			"error", err.Error(),
		)

		if err, ok := err.(*pq.Error); ok {
			switch err.Code {
			case "P0002", "02000":
				return nil, fmt.Errorf("%w, post not found", errlib.ErrNotFound)
			}
		}
		if err.Error() == "sql: no rows in result set" {
			return nil, fmt.Errorf("%w, post not found", errlib.ErrNotFound)
		}
		return nil, fmt.Errorf("%w, failed to get post", errlib.ErrInternal)
	}

	return &post, nil
}

func (r *psqlPostRepository) GetByIds(ctx context.Context, ids []uuid.UUID) ([]domain.Post, error) {
	posts := make([]domain.Post, 0)
	query := `select * from post where id IN (?)`

	query, args, err := sqlx.In(query, ids)
	if err != nil {
		slog.Error(
			"Some error while trying to get Posts by ids",
			"error", err.Error(),
		)

		return nil, fmt.Errorf("%w, failed to get posts", errlib.ErrInternal)
	}

	query = r.db.Rebind(query)
	if err := r.db.SelectContext(ctx, &posts, query, args...); err != nil {
		return nil, fmt.Errorf("%w, failed to fetch posts", errlib.ErrInternal)
	}

	return posts, nil
}

// GetRange gets first limit results after provided id (after)
//
// if after is nil, then returned first limit posts
func (r *psqlPostRepository) GetRange(ctx context.Context, after *uuid.UUID, limit uint) ([]domain.Post, error) {
	posts := make([]domain.Post, 0)

	var query string
	var err error

	if after != nil {
		query = `
			select * from post
			where id > $1
			order by id
			fetch first $2 row only
		`
		err = r.db.SelectContext(ctx, &posts, query, after, limit)
	} else {
		query = `
			select * from post
			order by id
			fetch first $1 row only
		`
		err = r.db.SelectContext(ctx, &posts, query, limit)
	}

	if err != nil {
		slog.Error(
			"Some error while trying to fetch Posts by range",
			"error", err.Error(),
		)

		return nil, fmt.Errorf("%w, failed to fetch posts", errlib.ErrInternal)
	}

	return posts, nil
}

func (r *psqlPostRepository) Count(ctx context.Context) (uint, error) {
	var count struct {
		Count uint `db:"cnt"`
	}

	if err := r.db.GetContext(ctx, &count, `select count(*) as cnt from post`); err != nil {
		slog.Error(
			"Some error while trying to count Posts",
			"error", err.Error(),
		)
		return 0, fmt.Errorf("%w, failed to get posts count", errlib.ErrInternal)
	}

	return count.Count, nil
}

func (r *psqlPostRepository) CreatePost(ctx context.Context, post domain.Post) (*domain.Post, error) {
	post.Id = uuid.New()
	query := `
		insert into post (id, title, authorId, createdAt, editedAt, content, commentsAllowed)
		values (:id, :title, :authorid, :createdat, :editedat, :content, :commentsallowed);
	`
	_, err := r.db.NamedExecContext(ctx, query, post)
	if err != nil {
		slog.Error(
			"Some error while trying to create Post",
			"error", err.Error(),
		)

		if err, ok := err.(*pq.Error); ok {
			switch err.Code {
			case "23505":
				return nil, fmt.Errorf("%w, failed to create post due to some key violation", errlib.ErrResourceAlreadyExists)
			}
		}

		return nil, fmt.Errorf("%w, failed to create user", errlib.ErrInternal)
	}

	return &post, nil
}

func (r *psqlPostRepository) UpdatePost(ctx context.Context, post domain.Post) (*domain.Post, error) {
	query := `
		update post
		set title=:title, authorid=:authorid, createdat=:createdat, editedat=:editedat, content=:content, commentsallowed=:commentsallowed
		where id=:id;
	`

	if _, err := r.db.NamedExecContext(ctx, query, post); err != nil {
		slog.Error(
			"Some error while trying to create Post",
			"error", err.Error(),
		)

		if err, ok := err.(*pq.Error); ok {
			switch err.Code {
			case "P0002", "02000":
				return nil, fmt.Errorf("%w, post not found", errlib.ErrNotFound)
			case "23505":
				return nil, fmt.Errorf("%w, failed to update post due to some key violation", errlib.ErrResourceAlreadyExists)
			}
		}

		return nil, fmt.Errorf("%w, failed to update post", errlib.ErrInternal)
	}
	return &post, nil
}
