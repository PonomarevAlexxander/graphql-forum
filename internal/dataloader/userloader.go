//go:generate go run github.com/vektah/dataloaden UserByCommentIdLoader DataLoaderByIdKey *github.com/PonomarevAlexxander/graphql-forum/internal/domain.User
//go:generate go run github.com/vektah/dataloaden UserByPostIdLoader DataLoaderByIdKey *github.com/PonomarevAlexxander/graphql-forum/internal/domain.User

package dataloader

import (
	"context"
	"time"

	"github.com/PonomarevAlexxander/graphql-forum/internal/domain"
	"github.com/PonomarevAlexxander/graphql-forum/internal/repository"
	"github.com/google/uuid"
)

func NewConfiguredUserByCommentIdLoader(repo repository.UserRepository, maxBatch int, waitTime time.Duration) *UserByCommentIdLoader {
	return NewUserByCommentIdLoader(UserByCommentIdLoaderConfig{
		Wait:     waitTime,
		MaxBatch: maxBatch,
		Fetch: func(keys []DataLoaderByIdKey) ([]*domain.User, []error) {
			items := make([]*domain.User, len(keys))
			errors := make([]error, len(keys))

			ids := uniqueIds(keys)
			users, err := repo.GetByCommentIds(context.Background(), ids)
			if err != nil {
				for index := range keys {
					errors[index] = err
				}
			}

			usersMap := groupUsersByCommentId(users)
			for i, key := range keys {
				if user, ok := usersMap[key.ID]; ok {
					items[i] = &user
				}
			}

			return items, errors
		},
	})
}

func NewConfiguredUserByPostIdLoader(repo repository.UserRepository, maxBatch int, waitTime time.Duration) *UserByPostIdLoader {
	return NewUserByPostIdLoader(UserByPostIdLoaderConfig{
		Wait:     waitTime,
		MaxBatch: maxBatch,
		Fetch: func(keys []DataLoaderByIdKey) ([]*domain.User, []error) {
			items := make([]*domain.User, len(keys))
			errors := make([]error, len(keys))

			ids := uniqueIds(keys)
			users, err := repo.GetByPostIds(context.Background(), ids)
			if err != nil {
				for index := range keys {
					errors[index] = err
				}
			}

			usersMap := groupUsersByPostId(users)
			for i, key := range keys {
				if user, ok := usersMap[key.ID]; ok {
					items[i] = &user
				}
			}

			return items, errors
		},
	})
}

func groupUsersByCommentId(users []domain.UserWithCommentId) map[uuid.UUID]domain.User {
	usersMap := make(map[uuid.UUID]domain.User)

	for _, user := range users {
		usersMap[user.CommentId] = user.User
	}

	return usersMap
}

func groupUsersByPostId(users []domain.UserWithPostId) map[uuid.UUID]domain.User {
	usersMap := make(map[uuid.UUID]domain.User)

	for _, user := range users {
		usersMap[user.PostId] = user.User
	}

	return usersMap
}
