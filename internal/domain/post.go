package domain

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	Id              uuid.UUID `db:"id"`
	Title           string    `db:"title"`
	AuthorId        uuid.UUID `db:"authorid"`
	CreatedAt       time.Time `db:"createdat"`
	EditedAt        time.Time `db:"editedat"`
	Content         string    `db:"content"`
	CommentsAllowed bool      `db:"commentsallowed"`
}
