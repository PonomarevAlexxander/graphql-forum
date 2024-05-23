//go:generate go run github.com/vektah/dataloaden CommentByIdLoader DataLoaderByIdKey *github.com/PonomarevAlexxander/graphql-forum/internal/domain.Comment
//go:generate go run github.com/vektah/dataloaden CommentByParentIdLoader DataLoaderByIdKey []github.com/PonomarevAlexxander/graphql-forum/internal/domain.Comment
//go:generate go run github.com/vektah/dataloaden CommentByPostIdLoader DataLoaderByIdKey []github.com/PonomarevAlexxander/graphql-forum/internal/domain.Comment

package dataloader

import (
	"context"
	"time"

	"github.com/PonomarevAlexxander/graphql-forum/internal/domain"
	"github.com/PonomarevAlexxander/graphql-forum/internal/repository"
	"github.com/google/uuid"
)

func NewConfiguredCommentByIdLoader(repo repository.CommentRepository, maxBatch int, waitTime time.Duration) *CommentByIdLoader {
	return NewCommentByIdLoader(CommentByIdLoaderConfig{
		Wait:     waitTime,
		MaxBatch: maxBatch,
		Fetch: func(keys []DataLoaderByIdKey) ([]*domain.Comment, []error) {
			items := make([]*domain.Comment, len(keys))
			errors := make([]error, len(keys))

			ids := uniqueIds(keys)
			comments, err := repo.GetByIds(context.Background(), ids)
			if err != nil {
				for index := range keys {
					errors[index] = err
				}
			}

			commentsMap := groupCommentsById(comments)
			for i, key := range keys {
				if comment, ok := commentsMap[key.ID]; ok {
					items[i] = &comment
				}
			}

			return items, errors
		},
	})
}

func NewConfiguredCommentByParentIdLoader(repo repository.CommentRepository, maxBatch int, waitTime time.Duration) *CommentByParentIdLoader {
	return NewCommentByParentIdLoader(CommentByParentIdLoaderConfig{
		Wait:     waitTime,
		MaxBatch: maxBatch,
		Fetch: func(keys []DataLoaderByIdKey) ([][]domain.Comment, []error) {
			items := make([][]domain.Comment, len(keys))
			errors := make([]error, len(keys))

			ids := uniqueIds(keys)
			comments, err := repo.GetRangeByParentIds(context.Background(), ids, 50)
			if err != nil {
				for index := range keys {
					errors[index] = err
				}
			}

			commentsMap := groupCommentsByParentId(comments)
			for i, key := range keys {
				if commentList, ok := commentsMap[key.ID]; ok {
					items[i] = commentList
				}
			}

			return items, errors
		},
	})
}

func NewConfiguredCommentByPostIdLoader(repo repository.CommentRepository, maxBatch int, waitTime time.Duration) *CommentByPostIdLoader {
	return NewCommentByPostIdLoader(CommentByPostIdLoaderConfig{
		Wait:     waitTime,
		MaxBatch: maxBatch,
		Fetch: func(keys []DataLoaderByIdKey) ([][]domain.Comment, []error) {
			items := make([][]domain.Comment, len(keys))
			errors := make([]error, len(keys))

			ids := uniqueIds(keys)
			comments, err := repo.GetRangeByPostIds(context.Background(), ids, 50)
			if err != nil {
				for index := range keys {
					errors[index] = err
				}
			}

			commentsMap := groupCommentsByPostId(comments)
			for i, key := range keys {
				if commentList, ok := commentsMap[key.ID]; ok {
					items[i] = commentList
				}
			}

			return items, errors
		},
	})
}

func groupCommentsById(comments []domain.Comment) map[uuid.UUID]domain.Comment {
	commentsMap := make(map[uuid.UUID]domain.Comment)

	for _, comment := range comments {
		commentsMap[comment.Id] = comment
	}

	return commentsMap
}

func groupCommentsByParentId(comments []domain.Comment) map[uuid.UUID][]domain.Comment {
	commentsMap := make(map[uuid.UUID][]domain.Comment)

	for _, comment := range comments {
		if comment.ParentId != nil {
			commentsMap[*comment.ParentId] = append(commentsMap[*comment.ParentId], comment)
		}
	}

	return commentsMap
}

func groupCommentsByPostId(comments []domain.Comment) map[uuid.UUID][]domain.Comment {
	commentsMap := make(map[uuid.UUID][]domain.Comment)

	for _, comment := range comments {
		if comment.PostId != nil {
			commentsMap[*comment.PostId] = append(commentsMap[*comment.PostId], comment)
		}
	}

	return commentsMap
}
