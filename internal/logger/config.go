package logger

type LoggerConfig struct {
	Level string `yaml:"level" validate:"required,oneof= 'info' 'error' 'warn' 'debug'"`
}
