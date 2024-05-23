package app

import (
	"github.com/PonomarevAlexxander/graphql-forum/internal/db"
	"github.com/PonomarevAlexxander/graphql-forum/internal/gql"
	"github.com/PonomarevAlexxander/graphql-forum/internal/logger"
	"github.com/PonomarevAlexxander/graphql-forum/internal/server"
)

type Config struct {
	DbConfig     db.DbConfig         `yaml:"db"`
	ServerConfig server.ServerConfig `yaml:"server"`
	LoggerConfig logger.LoggerConfig `yaml:"logger"`
	GqlConfig    gql.GqlConfig       `yaml:"gql"`
}
