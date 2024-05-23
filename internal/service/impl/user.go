package impl

import (
	"context"

	"github.com/PonomarevAlexxander/graphql-forum/internal/converter"
	"github.com/PonomarevAlexxander/graphql-forum/internal/dto"
	"github.com/PonomarevAlexxander/graphql-forum/internal/repository"
	"github.com/google/uuid"
)

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *userService {
	return &userService{
		repo: repo,
	}
}

func (s *userService) Create(ctx context.Context, user dto.UserCreateInputDto) (dto.UserDto, error) {
	created, err := s.repo.CreateUser(ctx, converter.DtoToDomainUserInput(user))
	if err != nil {
		return dto.UserDto{}, err
	}

	return converter.DomainToDtoUser(created), nil
}

func (s *userService) Get(ctx context.Context, id uuid.UUID) (dto.UserDto, error) {
	user, err := s.repo.GetById(ctx, id)
	if err != nil {
		return dto.UserDto{}, err
	}

	return converter.DomainToDtoUser(user), nil
}
