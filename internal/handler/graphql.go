package handler

import (
	"context"
	"fmt"
	"log/slog"

	gqlhandler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/PonomarevAlexxander/graphql-forum/internal/gql"
	"github.com/PonomarevAlexxander/graphql-forum/internal/gql/resolver"
	"github.com/PonomarevAlexxander/graphql-forum/internal/gql/runtime"
	"github.com/PonomarevAlexxander/graphql-forum/internal/service"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

const (
	maxUploadSize = 30 * 1024 * 1024
)

func newGraphQLHandler(services *service.Services, cfg *gql.GqlConfig) *gqlhandler.Server {
	handler := gqlhandler.New(
		runtime.NewExecutableSchema(
			newSchemaConfig(services),
		),
	)

	handler.AddTransport(transport.Websocket{
		KeepAlivePingInterval: cfg.WebSocketAlive,
	})
	handler.AddTransport(transport.Options{})
	handler.AddTransport(transport.POST{})
	handler.AddTransport(transport.MultipartForm{
		MaxUploadSize: maxUploadSize,
		MaxMemory:     maxUploadSize / 10,
	})

	handler.Use(extension.Introspection{})
	handler.Use(extension.FixedComplexityLimit(cfg.MaxComplexity))

	handler.SetRecoverFunc(func(ctx context.Context, err interface{}) error {
		slog.Error("unexpected error", "error", fmt.Sprintf("%v", err))
		return gqlerror.Errorf("internal server error")
	})

	return handler
}

func newSchemaConfig(services *service.Services) runtime.Config {
	cfg := runtime.Config{Resolvers: resolver.NewResolver(services)}

	cfg.Complexity.CommentsConnection.Edges = func(childComplexity int) int {
		return 2 * childComplexity
	}
	cfg.Complexity.CommentEdge.Replies = func(childComplexity int, first *uint, after *string) int {
		if first != nil {
			return childComplexity * int(*first)
		}
		return childComplexity * 2
	}
	cfg.Complexity.PostFindAllList.Edges = func(childComplexity int) int {
		return 2 * childComplexity
	}

	return cfg
}
