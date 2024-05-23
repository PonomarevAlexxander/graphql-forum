package db

type DbConfig struct {
	Type     string `yaml:type validate:"required,oneof='in-memory' 'psql'"`
	Host     string `yaml:"host" validate:"required"`
	Name     string `yaml:"name" validate:"required"`
	User     string `yaml:"user" validate:"required"`
	Password string `yaml:"password" validate:"required"`
}
