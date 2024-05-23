package repository

import (
	"fmt"

	"github.com/PonomarevAlexxander/graphql-forum/internal/db"
	"github.com/PonomarevAlexxander/graphql-forum/internal/errlib"
	"github.com/PonomarevAlexxander/graphql-forum/internal/repository/inmem"
	"github.com/PonomarevAlexxander/graphql-forum/internal/repository/psql"
)

type RepoType = string

const (
	PsqlRepositories     RepoType = "psql"
	InMemoryRepositories RepoType = "in-memory"
)

type Repositories struct {
	User    UserRepository
	Post    PostRepository
	Comment CommentRepository
}

func NewRepositories(repoType string, db *db.DbProvider) (*Repositories, error) {
	switch repoType {
	case PsqlRepositories:
		return &Repositories{
			User:    psql.NewPsqlUserRepository(db),
			Post:    psql.NewPsqlPostRepository(db),
			Comment: psql.NewPsqlCommentRepository(db),
		}, nil
	case InMemoryRepositories:
		memDb := inmem.NewDb()
		return &Repositories{
			User:    inmem.NewInMemUserRepository(memDb),
			Post:    inmem.NewInMemPostRepository(memDb),
			Comment: inmem.NewInMemCommentRepository(memDb),
		}, nil
	}
	return nil, fmt.Errorf("%w, unknown repository type passed %s", errlib.ErrInternal, repoType)
}
