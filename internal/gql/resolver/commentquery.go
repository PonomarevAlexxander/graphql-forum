package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.47

import (
	"context"

	"github.com/PonomarevAlexxander/graphql-forum/internal/gql/model"
	exec "github.com/PonomarevAlexxander/graphql-forum/internal/gql/runtime"
)

// Comment is the resolver for the comment field.
func (r *queryResolver) Comment(ctx context.Context) (*model.CommentQuery, error) {
	return &model.CommentQuery{}, nil
}

// CommentQuery returns exec.CommentQueryResolver implementation.
func (r *Resolver) CommentQuery() exec.CommentQueryResolver { return &commentQueryResolver{r} }

type commentQueryResolver struct{ *Resolver }