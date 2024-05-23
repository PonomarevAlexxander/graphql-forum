package dto

import (
	"time"

	"github.com/google/uuid"
)

type CommentDto struct {
	Id        uuid.UUID
	AuthorId  uuid.UUID
	PostId    *uuid.UUID
	ParentId  *uuid.UUID
	CreatedAt time.Time
	EditedAt  time.Time
	Content   string
}

type CommentCreateInput struct {
	PostID   uuid.UUID
	ParentID *uuid.UUID
	UserID   uuid.UUID
	Content  string
}

type CommentUpdateInput struct {
	UserID    uuid.UUID
	CommentID uuid.UUID
	Content   *string
}
