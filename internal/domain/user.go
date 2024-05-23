package domain

import "github.com/google/uuid"

type User struct {
	Id        uuid.UUID `db:"id"`
	Email     string    `db:"email"`
	FirstName string    `db:"firstname"`
	LastName  string    `db:"lastname"`
}

type UserWithCommentId struct {
	User
	CommentId uuid.UUID `db:"commentid"`
}

type UserWithPostId struct {
	User
	PostId uuid.UUID `db:"postid"`
}
