package dataloader

import (
	"github.com/PonomarevAlexxander/graphql-forum/internal/repository"
	"github.com/google/uuid"
)

type DataLoaderByIdKey struct {
	ID uuid.UUID
}

type DataLoaders struct {
	UserByCommentId   *UserByCommentIdLoader
	UserByPostId      *UserByPostIdLoader
	PostById          *PostByIdLoader
	CommentById       *CommentByIdLoader
	CommentByParentId *CommentByParentIdLoader
	CommentByPostId   *CommentByPostIdLoader
}

func NewDataLoaders(repos *repository.Repositories, config *DataLoadersConfig) *DataLoaders {
	return &DataLoaders{
		UserByCommentId:   NewConfiguredUserByCommentIdLoader(repos.User, config.MaxBatchSize, config.WaitTime),
		UserByPostId:      NewConfiguredUserByPostIdLoader(repos.User, config.MaxBatchSize, config.WaitTime),
		PostById:          NewConfiguredPostByIdLoader(repos.Post, config.MaxBatchSize, config.WaitTime),
		CommentById:       NewConfiguredCommentByIdLoader(repos.Comment, config.MaxBatchSize, config.WaitTime),
		CommentByParentId: NewConfiguredCommentByParentIdLoader(repos.Comment, config.MaxBatchSize, config.WaitTime),
		CommentByPostId:   NewConfiguredCommentByPostIdLoader(repos.Comment, config.MaxBatchSize, config.WaitTime),
	}
}

func uniqueIds(dataLoadIds []DataLoaderByIdKey) []uuid.UUID {
	mapping := make(map[uuid.UUID]bool)

	for _, key := range dataLoadIds {
		mapping[key.ID] = true
	}

	ids := make([]uuid.UUID, len(mapping))

	i := 0
	for key := range mapping {
		ids[i] = key
		i++
	}

	return ids
}
