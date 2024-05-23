package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.47

import (
	"context"

	"github.com/PonomarevAlexxander/graphql-forum/internal/gql/model"
	exec "github.com/PonomarevAlexxander/graphql-forum/internal/gql/runtime"
)

// Post is the resolver for the post field.
func (r *mutationResolver) Post(ctx context.Context) (*model.PostMutation, error) {
	return &model.PostMutation{}, nil
}

// PostMutation returns exec.PostMutationResolver implementation.
func (r *Resolver) PostMutation() exec.PostMutationResolver { return &postMutationResolver{r} }

type postMutationResolver struct{ *Resolver }
