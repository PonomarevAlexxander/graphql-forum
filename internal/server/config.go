package server

import "time"

type ServerConfig struct {
	Host         string        `yaml:"host" validate:"hostname_port,required"`
	ReadTimeout  time.Duration `yaml:"read-timeout" validate:"required"`
	WriteTimeout time.Duration `yaml:"write-timeout" validate:"required"`
}
