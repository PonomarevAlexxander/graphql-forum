package gql

import (
	"time"

	"github.com/PonomarevAlexxander/graphql-forum/internal/dataloader"
)

type GqlConfig struct {
	MaxComplexity  int                          `yaml:"max-complexity" validate:"required"`
	WebSocketAlive time.Duration                `yaml:"web-socket-live" validate:"required"`
	DataLoaders    dataloader.DataLoadersConfig `yaml:"data-loaders" validate:"required"`
}
