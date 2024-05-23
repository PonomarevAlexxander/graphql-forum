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

type psqlCommentRepository struct {
	db *db.DbProvider
}

func NewPsqlCommentRepository(db *db.DbProvider) *psqlCommentRepository {
	return &psqlCommentRepository{
		db: db,
	}
}

func (r *psqlCommentRepository) GetById(ctx context.Context, id uuid.UUID) (*domain.Comment, error) {
	var comment domain.Comment
	if err := r.db.GetContext(ctx, &comment, `select * from comment where id=$1`, id); err != nil {
		slog.Error(
			"Some error while trying to fetch Comment",
			"error", err.Error(),
		)

		if err, ok := err.(*pq.Error); ok {
			switch err.Code {
			case "P0002", "02000":
				return nil, fmt.Errorf("%w, comment not found", errlib.ErrNotFound)
			}
		}
		return nil, fmt.Errorf("%w, failed to get comment", errlib.ErrInternal)
	}

	return &comment, nil
}

func (r *psqlCommentRepository) GetByIds(ctx context.Context, ids []uuid.UUID) ([]domain.Comment, error) {
	comments := make([]domain.Comment, 0)
	query := `select * from comment where id IN (?)`

	query, args, err := sqlx.In(query, ids)
	if err != nil {
		slog.Error(
			"Some error while trying to fetch Comments",
			"error", err.Error(),
		)

		return nil, fmt.Errorf("%w, failed to get comments", errlib.ErrInternal)
	}

	query = r.db.Rebind(query)
	if err := r.db.SelectContext(ctx, &comments, query, args...); err != nil {
		return nil, fmt.Errorf("%w, failed to fetch comments", errlib.ErrInternal)
	}

	return comments, nil
}

func (r *psqlCommentRepository) CountByParentId(ctx context.Context, parentId uuid.UUID) (uint, error) {
	var count struct {
		Count uint `db:"cnt"`
	}

	if err := r.db.GetContext(ctx, &count, `select count(*) as cnt from comment where parentId=$1`, parentId); err != nil {
		slog.Error(
			"Some error while trying to count Comments by parent Id",
			"error", err.Error(),
		)

		if err, ok := err.(*pq.Error); ok {
			switch err.Code {
			case "P0002", "02000":
				return 0, fmt.Errorf("%w, post not found", errlib.ErrNotFound)
			}
		}
		return 0, fmt.Errorf("%w, failed to get comments count", errlib.ErrInternal)
	}

	return count.Count, nil
}

func (r *psqlCommentRepository) CountByPostId(ctx context.Context, postId uuid.UUID) (uint, error) {
	var count struct {
		Count uint `db:"cnt"`
	}

	if err := r.db.GetContext(ctx, &count, `select count(*) as cnt from comment where postId=$1`, postId); err != nil {
		slog.Error(
			"Some error while trying to count Comments by post Id",
			"error", err.Error(),
		)

		if err, ok := err.(*pq.Error); ok {
			switch err.Code {
			case "P0002", "02000":
				return 0, fmt.Errorf("%w, user not found", errlib.ErrNotFound)
			}
		}
		return 0, fmt.Errorf("%w, failed to get comments count", errlib.ErrInternal)
	}

	return count.Count, nil
}

// GetRangeByParentId gets first limit results after provided id (after)
func (r *psqlCommentRepository) GetRangeByParentId(ctx context.Context, after *uuid.UUID, limit uint, parentId uuid.UUID) ([]domain.Comment, error) {
	comments := make([]domain.Comment, 0)

	var query string
	var err error

	if after != nil {
		query = `
			select * from comment
			where parentId=$1 and id > $2
			order by id
			fetch first $3 row only
		`

		err = r.db.SelectContext(ctx, &comments, query, parentId, after, limit)
	} else {
		query = `
			select * from comment
			where parentId=$1
			order by id
			fetch first $2 row only
		`

		err = r.db.SelectContext(ctx, &comments, query, parentId, limit)
	}

	if err != nil {
		slog.Error(
			"Some error while trying to fetch Comments range by parent Id",
			"error", err.Error(),
		)

		if err, ok := err.(*pq.Error); ok {
			switch err.Code {
			case "P0002", "02000":
				return nil, fmt.Errorf("%w, parentId not found", errlib.ErrNotFound)
			}
		}

		return nil, fmt.Errorf("%w, failed to fetch comments", errlib.ErrInternal)
	}

	return comments, nil
}

// GetRangeByPostId gets first limit results after provided id (after)
func (r *psqlCommentRepository) GetRangeByPostId(ctx context.Context, after *uuid.UUID, limit uint, postId uuid.UUID) ([]domain.Comment, error) {
	comments := make([]domain.Comment, 0)

	var query string
	var err error

	if after != nil {
		query = `
			select * from comment
			where postId=$1 and id > $2
			order by id
			fetch first $3 row only
		`

		err = r.db.SelectContext(ctx, &comments, query, postId, after, limit)
	} else {
		query = `
			select * from comment
			where postId=$1
			order by id
			fetch first $2 row only
		`

		err = r.db.SelectContext(ctx, &comments, query, postId, limit)
	}

	if err != nil {
		slog.Error(
			"Some error while trying to fetch Comments range by post Id",
			"error", err.Error(),
		)

		if err, ok := err.(*pq.Error); ok {
			switch err.Code {
			case "P0002", "02000":
				return nil, fmt.Errorf("%w, postId not found", errlib.ErrNotFound)
			}
		}

		return nil, fmt.Errorf("%w, failed to fetch comments", errlib.ErrInternal)
	}

	return comments, nil
}

func (r *psqlCommentRepository) GetRangeByParentIds(ctx context.Context, parentIds []uuid.UUID, limitPerParent uint) ([]domain.Comment, error) {
	comments := make([]domain.Comment, 0)
	query := `
		select t.id, t.authorId, t.postId, t.parentId, t.createdAt, t.editedAt, t.content
		from (
			select c.*, row_number() over(
				partition by parentId
			) rn
			from comment c
			where parentId in (?)
		) t
		where rn <= ?
		order by id
	`
	query, args, err := sqlx.In(query, parentIds, limitPerParent)
	if err != nil {
		slog.Error(
			"Some error while trying to fetch Comments range by parent Ids",
			"error", err.Error(),
		)

		return nil, fmt.Errorf("%w, failed to get comments", errlib.ErrInternal)
	}
	query = r.db.Rebind(query)
	if err := r.db.SelectContext(ctx, &comments, query, args...); err != nil {
		slog.Error(
			"Some error while trying to fetch Comments range by parent Ids",
			"error", err.Error(),
		)

		if err, ok := err.(*pq.Error); ok {
			switch err.Code {
			case "P0002", "02000":
				return nil, fmt.Errorf("%w, parentId not found", errlib.ErrNotFound)
			}
		}

		return nil, fmt.Errorf("%w, failed to fetch comments", errlib.ErrInternal)
	}

	return comments, nil
}

func (r *psqlCommentRepository) GetRangeByPostIds(ctx context.Context, postIds []uuid.UUID, limitPerPost uint) ([]domain.Comment, error) {
	comments := make([]domain.Comment, 0)
	query := `
		select t.id, t.authorId, t.postId, t.parentId, t.createdAt, t.editedAt, t.content
		from (
			select c.*, row_number() over(
				partition by postId
			) rn
			from comment c
			where postId in (?)
		) t
		where rn <= ?
		order by id
	`
	query, args, err := sqlx.In(query, postIds, limitPerPost)
	if err != nil {
		slog.Error(
			"Some error while trying to fetch Comments range by post Ids",
			"error", err.Error(),
		)

		return nil, fmt.Errorf("%w, failed to get comments", errlib.ErrInternal)
	}
	query = r.db.Rebind(query)
	if err := r.db.SelectContext(ctx, &comments, query, args...); err != nil {
		slog.Error(
			"Some error while trying to fetch Comments range by post Ids",
			"error", err.Error(),
		)

		if err, ok := err.(*pq.Error); ok {
			switch err.Code {
			case "P0002", "02000":
				return nil, fmt.Errorf("%w, parentId not found", errlib.ErrNotFound)
			}
		}

		return nil, fmt.Errorf("%w, failed to fetch comments", errlib.ErrInternal)
	}

	return comments, nil
}

func (r *psqlCommentRepository) CreateComment(ctx context.Context, post domain.Comment) (*domain.Comment, error) {
	post.Id = uuid.New()
	query := `
		insert into comment (id, authorId, postId, parentId, createdAt, editedAt, content)
		values (:id, :authorid, :postid, :parentid, :createdat, :editedat, :content);
	`
	_, err := r.db.NamedExecContext(ctx, query, post)
	if err != nil {
		slog.Error(
			"Some error while trying to create Comment",
			"error", err.Error(),
		)

		if err, ok := err.(*pq.Error); ok {
			switch err.Code {
			case "23505":
				return nil, fmt.Errorf("%w, failed to create post due to some key violation", errlib.ErrResourceAlreadyExists)
			}
		}

		return nil, fmt.Errorf("%w, failed to create post", errlib.ErrInternal)
	}

	return &post, nil
}

func (r *psqlCommentRepository) UpdateComment(ctx context.Context, comment domain.Comment) (*domain.Comment, error) {
	query := `
		update comment
		set authorId=:authorid, postId=:postid, parentId=:parentid, createdAt=:createdat, editedAt=:editedat, content=:content
		where id=:id;
	`

	if _, err := r.db.NamedExecContext(ctx, query, comment); err != nil {
		slog.Error(
			"Some error while trying to update Comment",
			"error", err.Error(),
		)

		if err, ok := err.(*pq.Error); ok {
			switch err.Code {
			case "P0002", "02000":
				return nil, fmt.Errorf("%w, comment not found", errlib.ErrNotFound)
			case "23505":
				return nil, fmt.Errorf("%w, failed to update comment due to some key violation", errlib.ErrResourceAlreadyExists)
			}
		}

		return nil, fmt.Errorf("%w, failed to update comment", errlib.ErrInternal)
	}
	return &comment, nil
}
