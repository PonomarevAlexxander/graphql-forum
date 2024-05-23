package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.47

import (
	"context"
	"errors"
	"log/slog"

	"github.com/99designs/gqlgen/graphql"
	"github.com/PonomarevAlexxander/graphql-forum/internal/config"
	"github.com/PonomarevAlexxander/graphql-forum/internal/converter"
	"github.com/PonomarevAlexxander/graphql-forum/internal/dataloader"
	"github.com/PonomarevAlexxander/graphql-forum/internal/dto"
	"github.com/PonomarevAlexxander/graphql-forum/internal/errlib"
	"github.com/PonomarevAlexxander/graphql-forum/internal/gql/model"
	exec "github.com/PonomarevAlexxander/graphql-forum/internal/gql/runtime"
	"github.com/google/uuid"
)

// TotalCount is the resolver for the totalCount field.
func (r *postFindAllListResolver) TotalCount(ctx context.Context, obj *model.PostFindAllList) (model.TotalCountResolvingResult, error) {
	count, err := r.s.Post.Count(ctx)
	if err != nil {
		slog.Error(
			"Some error while trying to fetch TotalCount from post",
			"error", err.Error(),
		)

		return model.InternalError{Message: err.Error()}, nil
	}

	return model.TotalCount{
		Value: count,
	}, nil
}

// Comments is the resolver for the comments field.
func (r *postFindElementResolver) Comments(ctx context.Context, obj *model.PostFindElement, first *uint, after *string) (model.CommentsConnectionResolvingResult, error) {
	slog.Debug("Fetching Comments...")
	dataLoader := ctx.Value(config.DataLoadersContextKey).(*dataloader.DataLoaders).CommentByPostId

	postId, ok := graphql.GetFieldContext(ctx).Parent.Args["postId"].(uuid.UUID)
	if !ok {
		slog.Error("Some error while trying to fetch Comments", "error", "invalid type")
		return model.InternalError{Message: errlib.ErrInternal.Error()}, nil
	}

	comments, err := dataLoader.Load(dataloader.DataLoaderByIdKey{ID: postId})
	if err != nil {
		slog.Error(
			"Some error while trying to fetch Comments",
			"error", err.Error(),
		)
		if errors.Is(err, errlib.ErrNotFound) {
			return model.NotFoundError{Message: err.Error()}, nil
		}
		return model.InternalError{Message: err.Error()}, nil
	}

	var connection model.CommentsConnection

	connection.Edges = converter.DomainCommentsToGqlEdges(comments)
	connection.PageInfo = converter.DomainCommentsToGqlPageInfo(comments)

	return connection, nil
}

// FindAll is the resolver for the findAll field.
func (r *postQueryResolver) FindAll(ctx context.Context, obj *model.PostQuery, first *uint, after *string) (model.PostFindAllResult, error) {
	var id *uuid.UUID = nil

	if after != nil {
		parsedId, err := uuid.Parse(*after)
		if err != nil {
			slog.Error(
				"Failed to parse UUID",
				"error", err.Error(),
			)

			return model.InternalError{Message: errlib.ErrInternal.Error()}, nil
		}

		id = &parsedId
	}

	posts, err := r.s.Post.GetRange(ctx, dto.PostGetDto{
		Limit: *first,
		After: id,
	})
	if err != nil {
		slog.Error(
			"Some error while trying to fetch all Posts",
			"error", err.Error(),
		)
		return model.InternalError{Message: err.Error()}, nil
	}

	return model.PostFindAllList{
		Edges:    converter.DtoPostsToGqlEdges(posts),
		PageInfo: converter.DtoPostsToGqlPageInfo(posts),
	}, nil
}

// Find is the resolver for the find field.
func (r *postQueryResolver) Find(ctx context.Context, obj *model.PostQuery, postID uuid.UUID) (model.PostFindResult, error) {
	post, err := r.s.Post.Get(ctx, postID)

	if err != nil {
		slog.Error(
			"Some error while trying find Post",
			"error", err.Error(),
		)

		if errors.Is(err, errlib.ErrNotFound) {
			return model.NotFoundError{Message: err.Error()}, nil
		}
		return model.InternalError{Message: err.Error()}, nil
	}

	return model.PostFindElement{
		Post: converter.DtoToGqlPost(post),
	}, nil
}

// PostFindAllList returns exec.PostFindAllListResolver implementation.
func (r *Resolver) PostFindAllList() exec.PostFindAllListResolver { return &postFindAllListResolver{r} }

// PostFindElement returns exec.PostFindElementResolver implementation.
func (r *Resolver) PostFindElement() exec.PostFindElementResolver { return &postFindElementResolver{r} }

type postFindAllListResolver struct{ *Resolver }
type postFindElementResolver struct{ *Resolver }