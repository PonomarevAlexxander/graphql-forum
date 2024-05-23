package dto

import "github.com/google/uuid"

type UserCreateInputDto struct {
	Email     string
	FirstName string
	LastName  string
}

type UserDto struct {
	Id        uuid.UUID
	Email     string
	FirstName string
	LastName  string
}
