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

type psqlUserRepository struct {
	db *db.DbProvider
}

func NewPsqlUserRepository(db *db.DbProvider) *psqlUserRepository {
	return &psqlUserRepository{
		db: db,
	}
}

func (r *psqlUserRepository) GetById(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	var user domain.User
	if err := r.db.GetContext(ctx, &user, `select * from "user" where id=$1`, id); err != nil {
		if err, ok := err.(*pq.Error); ok {
			switch err.Code {
			case "P0002", "02000":
				return nil, fmt.Errorf("%w, user not found", errlib.ErrNotFound)
			}
		}
		return nil, fmt.Errorf("%w, failed to get user", errlib.ErrInternal)
	}

	return &user, nil
}

func (r *psqlUserRepository) GetByIds(ctx context.Context, ids []uuid.UUID) ([]domain.User, error) {
	users := make([]domain.User, 0)
	query := `select * from "user" where id IN (?)`

	query, args, err := sqlx.In(query, ids)
	if err != nil {
		return nil, fmt.Errorf("%w, failed to get users", errlib.ErrInternal)
	}

	query = r.db.Rebind(query)
	if err := r.db.SelectContext(ctx, &users, query, args...); err != nil {
		return nil, fmt.Errorf("%w, failed to fetch users", errlib.ErrInternal)
	}

	return users, nil
}

func (r *psqlUserRepository) GetByCommentIds(ctx context.Context, ids []uuid.UUID) ([]domain.UserWithCommentId, error) {
	users := make([]domain.UserWithCommentId, 0)
	query := `
		select u.*, c.id as commentId
		from "user" u
		join comment c
		on c.authorId = u.id
		where c.id IN (?)
	`

	query, args, err := sqlx.In(query, ids)
	if err != nil {
		slog.Error(
			"Some error while trying to fetch Users by Comment ids",
			"error", err.Error(),
		)

		return nil, fmt.Errorf("%w, failed to get users", errlib.ErrInternal)
	}

	query = r.db.Rebind(query)
	if err := r.db.SelectContext(ctx, &users, query, args...); err != nil {
		slog.Error(
			"Some error while trying to fetch Users by Comment ids",
			"error", err.Error(),
		)

		return nil, fmt.Errorf("%w, failed to fetch users", errlib.ErrInternal)
	}

	return users, nil
}

func (r *psqlUserRepository) GetByPostIds(ctx context.Context, ids []uuid.UUID) ([]domain.UserWithPostId, error) {
	users := make([]domain.UserWithPostId, 0)
	query := `
	select u.*, p.id as postId
	from "user" u
	join post p
	on p.authorId = u.id
	where p.id IN (?)
`

	query, args, err := sqlx.In(query, ids)
	if err != nil {
		slog.Error(
			"Some error while trying to fetch Users by Post ids",
			"error", err.Error(),
		)

		return nil, fmt.Errorf("%w, failed to get users", errlib.ErrInternal)
	}

	query = r.db.Rebind(query)
	if err := r.db.SelectContext(ctx, &users, query, args...); err != nil {
		slog.Error(
			"Some error while trying to fetch Users by Post ids",
			"error", err.Error(),
			"query", query,
		)
		return nil, fmt.Errorf("%w, failed to fetch users", errlib.ErrInternal)
	}

	return users, nil
}

func (r *psqlUserRepository) CreateUser(ctx context.Context, user domain.User) (*domain.User, error) {
	user.Id = uuid.New()
	query := `
		insert into "user" (id, email, firstName, lastName)
		values (:id, :email, :firstname, :lastname);
	`
	_, err := r.db.NamedExecContext(ctx, query, user)
	if err != nil {
		slog.Error(
			"Some error while trying to create User",
			"error", err.Error(),
		)

		if err, ok := err.(*pq.Error); ok {
			switch err.Code {
			case "23505":
				return nil, fmt.Errorf("%w, failed to create user due to some key violation", errlib.ErrResourceAlreadyExists)
			}
		}

		return nil, fmt.Errorf("%w, failed to create user", errlib.ErrInternal)
	}

	return &user, nil
}
