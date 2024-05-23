package domain

import (
	"time"

	"github.com/google/uuid"
)

type Comment struct {
	Id        uuid.UUID  `db:"id"`
	AuthorId  uuid.UUID  `db:"authorid"`
	PostId    *uuid.UUID `db:"postid"`
	ParentId  *uuid.UUID `db:"parentid"`
	CreatedAt time.Time  `db:"createdat"`
	EditedAt  time.Time  `db:"editedat"`
	Content   string     `db:"content"`
}
