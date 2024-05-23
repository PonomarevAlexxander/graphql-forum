package dto

import (
	"time"

	"github.com/google/uuid"
)

type PostCreateDto struct {
	UserId          uuid.UUID
	Title           string
	Content         string
	CommentsAllowed bool
}

type PostUpdateDto struct {
	UserId          uuid.UUID
	PostId          uuid.UUID
	Title           *string
	Content         *string
	CommentsAllowed *bool
}

type PostGetDto struct {
	Limit uint
	After *uuid.UUID
}

type PostsListDto struct {
	Posts []PostDto
}

type PostDto struct {
	Id              uuid.UUID
	Title           string
	UserID          uuid.UUID
	CreatedAt       time.Time
	EditedAt        time.Time
	Content         string
	CommentsAllowed bool
}
