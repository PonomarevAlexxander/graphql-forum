// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"

	"github.com/google/uuid"
)

type CommentCreateResult interface {
	IsCommentCreateResult()
}

type CommentFindResult interface {
	IsCommentFindResult()
}

type CommentResolvingResult interface {
	IsCommentResolvingResult()
}

type CommentUpdateResult interface {
	IsCommentUpdateResult()
}

type CommentsConnectionResolvingResult interface {
	IsCommentsConnectionResolvingResult()
}

type ErrorInterface interface {
	IsErrorInterface()
	GetMessage() string
}

type PostCreateResult interface {
	IsPostCreateResult()
}

type PostFindAllResult interface {
	IsPostFindAllResult()
}

type PostFindResult interface {
	IsPostFindResult()
}

type PostResolvingResult interface {
	IsPostResolvingResult()
}

type PostUpdateResult interface {
	IsPostUpdateResult()
}

type TotalCountResolvingResult interface {
	IsTotalCountResolvingResult()
}

type UserCreateResult interface {
	IsUserCreateResult()
}

type UserResolvingResult interface {
	IsUserResolvingResult()
}

type BadRequestError struct {
	Message string `json:"message"`
}

func (BadRequestError) IsCommentCreateResult() {}

func (BadRequestError) IsErrorInterface()       {}
func (this BadRequestError) GetMessage() string { return this.Message }

type Comment struct {
	ID        uuid.UUID           `json:"id"`
	ParentID  *uuid.UUID          `json:"parentId,omitempty"`
	PostID    *uuid.UUID          `json:"postId,omitempty"`
	Author    UserResolvingResult `json:"author,omitempty"`
	CreatedAt *time.Time          `json:"createdAt,omitempty"`
	EditedAt  *time.Time          `json:"editedAt,omitempty"`
	Content   *string             `json:"content,omitempty"`
}

func (Comment) IsCommentResolvingResult() {}

type CommentCreateInput struct {
	PostID   uuid.UUID  `json:"postId"`
	ParentID *uuid.UUID `json:"parentId,omitempty"`
	UserID   uuid.UUID  `json:"userId"`
	Content  string     `json:"content"`
}

type CommentCreateOk struct {
	Comment *Comment `json:"comment"`
}

func (CommentCreateOk) IsCommentCreateResult() {}

type CommentEdge struct {
	Comment CommentResolvingResult            `json:"comment"`
	Cursor  *string                           `json:"cursor,omitempty"`
	Replies CommentsConnectionResolvingResult `json:"replies,omitempty"`
}

type CommentFindElement struct {
	Comment CommentResolvingResult            `json:"comment"`
	Replies CommentsConnectionResolvingResult `json:"replies,omitempty"`
}

func (CommentFindElement) IsCommentFindResult() {}

type CommentMutation struct {
	Create CommentCreateResult `json:"create"`
	Update CommentUpdateResult `json:"update"`
}

type CommentQuery struct {
	Find CommentFindResult `json:"find"`
}

type CommentUpdateInput struct {
	UserID    uuid.UUID `json:"userId"`
	CommentID uuid.UUID `json:"commentId"`
	Content   *string   `json:"content,omitempty"`
}

type CommentUpdateOk struct {
	Comment *Comment `json:"comment"`
}

func (CommentUpdateOk) IsCommentUpdateResult() {}

type CommentsConnection struct {
	TotalCount TotalCountResolvingResult `json:"totalCount,omitempty"`
	Edges      []*CommentEdge            `json:"edges,omitempty"`
	PageInfo   *PageInfo                 `json:"pageInfo,omitempty"`
}

func (CommentsConnection) IsCommentsConnectionResolvingResult() {}

type ConflictError struct {
	Message string `json:"message"`
}

func (ConflictError) IsCommentCreateResult() {}

func (ConflictError) IsCommentUpdateResult() {}

func (ConflictError) IsUserCreateResult() {}

func (ConflictError) IsErrorInterface()       {}
func (this ConflictError) GetMessage() string { return this.Message }

type InternalError struct {
	Message string `json:"message"`
}

func (InternalError) IsCommentCreateResult() {}

func (InternalError) IsCommentUpdateResult() {}

func (InternalError) IsPostCreateResult() {}

func (InternalError) IsPostUpdateResult() {}

func (InternalError) IsUserCreateResult() {}

func (InternalError) IsCommentFindResult() {}

func (InternalError) IsPostFindAllResult() {}

func (InternalError) IsPostFindResult() {}

func (InternalError) IsErrorInterface()       {}
func (this InternalError) GetMessage() string { return this.Message }

func (InternalError) IsCommentsConnectionResolvingResult() {}

func (InternalError) IsCommentResolvingResult() {}

func (InternalError) IsPostResolvingResult() {}

func (InternalError) IsTotalCountResolvingResult() {}

func (InternalError) IsUserResolvingResult() {}

type Mutation struct {
}

type NotFoundError struct {
	Message string `json:"message"`
}

func (NotFoundError) IsCommentCreateResult() {}

func (NotFoundError) IsCommentUpdateResult() {}

func (NotFoundError) IsPostCreateResult() {}

func (NotFoundError) IsPostUpdateResult() {}

func (NotFoundError) IsCommentFindResult() {}

func (NotFoundError) IsPostFindAllResult() {}

func (NotFoundError) IsPostFindResult() {}

func (NotFoundError) IsErrorInterface()       {}
func (this NotFoundError) GetMessage() string { return this.Message }

func (NotFoundError) IsCommentsConnectionResolvingResult() {}

func (NotFoundError) IsCommentResolvingResult() {}

func (NotFoundError) IsPostResolvingResult() {}

func (NotFoundError) IsUserResolvingResult() {}

type PageInfo struct {
	HasNextPage *bool   `json:"hasNextPage,omitempty"`
	EndCursor   *string `json:"endCursor,omitempty"`
}

type Post struct {
	ID              uuid.UUID           `json:"id"`
	Title           *string             `json:"title,omitempty"`
	Author          UserResolvingResult `json:"author,omitempty"`
	CreatedAt       *time.Time          `json:"createdAt,omitempty"`
	EditedAt        *time.Time          `json:"editedAt,omitempty"`
	Content         *string             `json:"content,omitempty"`
	CommentsAllowed *bool               `json:"commentsAllowed,omitempty"`
}

func (Post) IsPostResolvingResult() {}

type PostCreateInput struct {
	UserID          uuid.UUID `json:"userId"`
	Title           string    `json:"title"`
	Content         string    `json:"content"`
	CommentsAllowed bool      `json:"commentsAllowed"`
}

type PostCreateOk struct {
	Post *Post `json:"post"`
}

func (PostCreateOk) IsPostCreateResult() {}

type PostEdge struct {
	Post   PostResolvingResult `json:"post"`
	Cursor *string             `json:"cursor,omitempty"`
}

type PostFindAllList struct {
	TotalCount TotalCountResolvingResult `json:"totalCount,omitempty"`
	Edges      []*PostEdge               `json:"edges,omitempty"`
	PageInfo   *PageInfo                 `json:"pageInfo,omitempty"`
}

func (PostFindAllList) IsPostFindAllResult() {}

type PostFindElement struct {
	Post     PostResolvingResult               `json:"post"`
	Comments CommentsConnectionResolvingResult `json:"comments,omitempty"`
}

func (PostFindElement) IsPostFindResult() {}

type PostMutation struct {
	Create PostCreateResult `json:"create"`
	Update PostUpdateResult `json:"update"`
}

type PostQuery struct {
	FindAll PostFindAllResult `json:"findAll"`
	Find    PostFindResult    `json:"find"`
}

type PostUpdateInput struct {
	UserID          uuid.UUID `json:"userId"`
	PostID          uuid.UUID `json:"postId"`
	Title           *string   `json:"title,omitempty"`
	Content         *string   `json:"content,omitempty"`
	CommentsAllowed *bool     `json:"commentsAllowed,omitempty"`
}

type PostUpdateOk struct {
	Post *Post `json:"post"`
}

func (PostUpdateOk) IsPostUpdateResult() {}

type Query struct {
}

type TotalCount struct {
	Value uint `json:"value"`
}

func (TotalCount) IsTotalCountResolvingResult() {}

type User struct {
	ID        uuid.UUID `json:"id"`
	Email     *string   `json:"email,omitempty"`
	FirstName *string   `json:"firstName,omitempty"`
	LastName  *string   `json:"lastName,omitempty"`
}

func (User) IsUserResolvingResult() {}

type UserCreateInput struct {
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type UserCreateOk struct {
	User *User `json:"user"`
}

func (UserCreateOk) IsUserCreateResult() {}

type UserMutation struct {
	Create UserCreateResult `json:"create"`
}
