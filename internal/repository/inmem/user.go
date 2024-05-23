package inmem

import (
	"context"

	"github.com/PonomarevAlexxander/graphql-forum/internal/domain"
	"github.com/google/uuid"
)

type inmemUserRepository struct {
	db *db
}

func NewInMemUserRepository(db *db) *inmemUserRepository {
	return &inmemUserRepository{
		db: db,
	}
}

func (r *inmemUserRepository) GetById(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	user, err := r.db.GetUser(id)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *inmemUserRepository) GetByIds(ctx context.Context, ids []uuid.UUID) ([]domain.User, error) {
	users := make([]domain.User, len(ids))
	var err error

	for i, id := range ids {
		users[i], err = r.db.GetUser(id)

		if err != nil {
			return nil, err
		}
	}

	return users, nil
}

func (r *inmemUserRepository) GetByCommentIds(ctx context.Context, ids []uuid.UUID) ([]domain.UserWithCommentId, error) {
	users := make([]domain.UserWithCommentId, 0)

	for _, id := range ids {
		user, err := r.db.GetUserByCommentId(id)
		if err != nil {
			return nil, err
		}

		users = append(users, domain.UserWithCommentId{
			User:      user,
			CommentId: id,
		})
	}

	return users, nil
}

func (r *inmemUserRepository) GetByPostIds(ctx context.Context, ids []uuid.UUID) ([]domain.UserWithPostId, error) {
	users := make([]domain.UserWithPostId, 0)

	for _, id := range ids {
		user, err := r.db.GetUserByPostId(id)
		if err != nil {
			return nil, err
		}

		users = append(users, domain.UserWithPostId{
			User:   user,
			PostId: id,
		})
	}

	return users, nil
}

func (r *inmemUserRepository) CreateUser(ctx context.Context, user domain.User) (*domain.User, error) {
	user.Id = uuid.New()
	err := r.db.NewUser(user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
