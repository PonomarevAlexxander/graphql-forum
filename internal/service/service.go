package service

import (
	"github.com/PonomarevAlexxander/graphql-forum/internal/repository"
	"github.com/PonomarevAlexxander/graphql-forum/internal/service/impl"
)

type Services struct {
	User    UserService
	Post    PostService
	Comment CommentService
}

func NewServices(repos *repository.Repositories) *Services {
	return &Services{
		User:    impl.NewUserService(repos.User),
		Post:    impl.NewPostService(repos.Post),
		Comment: impl.NewCommentService(repos),
	}
}
