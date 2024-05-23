package handler

import (
	"context"
	"net/http"

	"github.com/PonomarevAlexxander/graphql-forum/internal/config"
	"github.com/PonomarevAlexxander/graphql-forum/internal/dataloader"
	"github.com/PonomarevAlexxander/graphql-forum/internal/repository"
)

func NewDataLoadersInjector(repos *repository.Repositories, cfg *dataloader.DataLoadersConfig) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			ctx := context.WithValue(
				r.Context(),
				config.DataLoadersContextKey,
				dataloader.NewDataLoaders(repos, cfg),
			)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
