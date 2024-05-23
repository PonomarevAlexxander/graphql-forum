package handler

import (
	"net/http"

	"github.com/PonomarevAlexxander/graphql-forum/internal/gql"
	"github.com/PonomarevAlexxander/graphql-forum/internal/repository"
	"github.com/PonomarevAlexxander/graphql-forum/internal/service"
)

func SetHandler(router *http.ServeMux, repos *repository.Repositories, services *service.Services, cfg *gql.GqlConfig) {
	dataLoader := NewDataLoadersInjector(repos, &cfg.DataLoaders)
	router.Handle("/graphql", dataLoader(newGraphQLHandler(services, cfg)))
}
