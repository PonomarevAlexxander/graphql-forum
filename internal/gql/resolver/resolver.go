package resolver

import "github.com/PonomarevAlexxander/graphql-forum/internal/service"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	s *service.Services
}

func NewResolver(services *service.Services) *Resolver {
	return &Resolver{
		s: services,
	}
}
