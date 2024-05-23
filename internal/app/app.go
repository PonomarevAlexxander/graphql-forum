package app

import (
	"fmt"
	"net/http"

	"github.com/PonomarevAlexxander/graphql-forum/internal/db"
	"github.com/PonomarevAlexxander/graphql-forum/internal/handler"
	"github.com/PonomarevAlexxander/graphql-forum/internal/repository"
	"github.com/PonomarevAlexxander/graphql-forum/internal/server"
	"github.com/PonomarevAlexxander/graphql-forum/internal/service"
	"github.com/rs/cors"
)

type App struct {
	server   *server.Server
	provider *db.DbProvider
}

func NewApp(config *Config, notify chan error) (*App, error) {
	provider, err := db.NewPsqlProvider(&config.DbConfig)
	if err != nil {
		return nil, fmt.Errorf("Failed to initialize db with error: %w", err)
	}

	repos, err := repository.NewRepositories(config.DbConfig.Type, provider)
	if err != nil {
		return nil, err
	}
	services := service.NewServices(repos)

	c := cors.New(cors.Options{
		AllowCredentials: true,
		AllowedMethods:   []string{http.MethodOptions, http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPut, http.MethodPatch},
		AllowedHeaders:   []string{"Content-Type", "Authorization", "Set-Cookie", "User-Agent", "Origin"},
		AllowOriginFunc: func(origin string) bool {
			return true
		},
	})

	router := http.NewServeMux()
	handler.SetHandler(router, repos, services, &config.GqlConfig)

	server := server.NewServer(&config.ServerConfig, c.Handler(router), notify)

	return &App{
		server:   server,
		provider: provider,
	}, nil
}

func (app *App) Start() {
	app.server.Start()
}

func (app *App) Stop() error {
	serverErr := app.server.Stop()
	providerErr := app.provider.Close()
	if serverErr != nil || providerErr != nil {
		return fmt.Errorf("Provider error: %w. Server error: %w", providerErr, serverErr)
	}
	return nil
}
