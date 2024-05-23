//go:generate go run github.com/vektah/dataloaden PostByIdLoader DataLoaderByIdKey *github.com/PonomarevAlexxander/graphql-forum/internal/domain.Post

package dataloader

import (
	"context"
	"time"

	"github.com/PonomarevAlexxander/graphql-forum/internal/domain"
	"github.com/PonomarevAlexxander/graphql-forum/internal/repository"
	"github.com/google/uuid"
)

func NewConfiguredPostByIdLoader(repo repository.PostRepository, maxBatch int, waitTime time.Duration) *PostByIdLoader {
	return NewPostByIdLoader(PostByIdLoaderConfig{
		Wait:     waitTime,
		MaxBatch: maxBatch,
		Fetch: func(keys []DataLoaderByIdKey) ([]*domain.Post, []error) {
			items := make([]*domain.Post, len(keys))
			errors := make([]error, len(keys))

			ids := uniqueIds(keys)
			posts, err := repo.GetByIds(context.Background(), ids)
			if err != nil {
				for index := range keys {
					errors[index] = err
				}
			}

			postsMap := groupPostsById(posts)
			for i, key := range keys {
				if post, ok := postsMap[key.ID]; ok {
					items[i] = &post
				}
			}

			return items, errors
		},
	})
}

func groupPostsById(posts []domain.Post) map[uuid.UUID]domain.Post {
	postsMap := make(map[uuid.UUID]domain.Post)

	for _, post := range posts {
		postsMap[post.Id] = post
	}

	return postsMap
}
