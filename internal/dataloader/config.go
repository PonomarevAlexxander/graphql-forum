package dataloader

import "time"

type DataLoadersConfig struct {
	WaitTime     time.Duration `yaml:"wait-time" validate:"required"`
	MaxBatchSize int           `yaml:"max-batch-size" validate:"required"`
}
