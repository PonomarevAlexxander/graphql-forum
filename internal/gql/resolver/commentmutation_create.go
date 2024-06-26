package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.47

import (
	"context"
	"errors"
	"log/slog"

	"github.com/PonomarevAlexxander/graphql-forum/internal/converter"
	"github.com/PonomarevAlexxander/graphql-forum/internal/errlib"
	"github.com/PonomarevAlexxander/graphql-forum/internal/gql/model"
)

// Create is the resolver for the create field.
func (r *commentMutationResolver) Create(ctx context.Context, obj *model.CommentMutation, input model.CommentCreateInput) (model.CommentCreateResult, error) {
	comment, err := r.s.Comment.Create(ctx, *converter.GqlToDtoCommentInput(input))

	if err != nil {
		slog.Error(
			"Some error while trying to create Comment",
			"error", err.Error(),
		)
		if errors.Is(err, errlib.ErrResourceAlreadyExists) {
			return model.ConflictError{Message: err.Error()}, nil
		}
		if errors.Is(err, errlib.ErrNotFound) {
			return model.NotFoundError{Message: err.Error()}, nil
		}
		return model.InternalError{Message: err.Error()}, nil
	}

	return model.CommentCreateOk{
		Comment: converter.DtoToGqlComment(comment),
	}, nil
}
