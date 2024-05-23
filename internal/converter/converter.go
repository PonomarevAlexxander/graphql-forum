package converter

import (
	"time"

	"github.com/PonomarevAlexxander/graphql-forum/internal/domain"
	"github.com/PonomarevAlexxander/graphql-forum/internal/dto"
	"github.com/PonomarevAlexxander/graphql-forum/internal/gql/model"
)

func DomainToGqlPost(post *domain.Post) model.Post {
	return model.Post{
		ID:              post.Id,
		Title:           &post.Title,
		CreatedAt:       &post.CreatedAt,
		EditedAt:        &post.EditedAt,
		Content:         &post.Content,
		CommentsAllowed: &post.CommentsAllowed,
	}
}

func DomainToDtoPost(post *domain.Post) dto.PostDto {
	return dto.PostDto{
		Id:              post.Id,
		Title:           post.Title,
		CreatedAt:       post.CreatedAt,
		EditedAt:        post.EditedAt,
		Content:         post.Content,
		CommentsAllowed: post.CommentsAllowed,
	}
}

func DomainToDtoPosts(posts []domain.Post) []dto.PostDto {
	dtoPosts := make([]dto.PostDto, len(posts))

	for i, post := range posts {
		dtoPosts[i] = DomainToDtoPost(&post)
	}

	return dtoPosts
}

func DtoToDomainInputPost(post dto.PostCreateDto) domain.Post {
	return domain.Post{
		AuthorId:        post.UserId,
		Title:           post.Title,
		Content:         post.Content,
		CommentsAllowed: post.CommentsAllowed,
		CreatedAt:       time.Now().UTC(),
		EditedAt:        time.Now().UTC(),
	}
}

func DtoToDomainUserInput(dto dto.UserCreateInputDto) domain.User {
	return domain.User{
		Email:     dto.Email,
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
	}
}

func DomainToDtoUser(user *domain.User) dto.UserDto {
	return dto.UserDto{
		Id:        user.Id,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}
}

func DtoToGqlPost(post dto.PostDto) *model.Post {
	return &model.Post{
		ID:              post.Id,
		Title:           &post.Title,
		CreatedAt:       &post.CreatedAt,
		EditedAt:        &post.EditedAt,
		Content:         &post.Content,
		CommentsAllowed: &post.CommentsAllowed,
	}
}

func DomainToGqlUser(user *domain.User) model.User {
	return model.User{
		ID:        user.Id,
		Email:     &user.Email,
		FirstName: &user.FirstName,
		LastName:  &user.LastName,
	}
}

func DtoToGqlUser(user dto.UserDto) *model.User {
	return &model.User{
		ID:        user.Id,
		Email:     &user.Email,
		FirstName: &user.FirstName,
		LastName:  &user.LastName,
	}
}

func GqlToDtoUserInput(userInput model.UserCreateInput) *dto.UserCreateInputDto {
	return &dto.UserCreateInputDto{
		Email:     userInput.Email,
		FirstName: userInput.FirstName,
		LastName:  userInput.LastName,
	}
}

func DomainToGqlComment(comment domain.Comment) *model.Comment {
	return &model.Comment{
		ID:        comment.Id,
		ParentID:  comment.ParentId,
		PostID:    comment.PostId,
		CreatedAt: &comment.CreatedAt,
		EditedAt:  &comment.EditedAt,
		Content:   &comment.Content,
	}
}

func DtoToDomainCommentInput(comment dto.CommentCreateInput) domain.Comment {
	dComment := domain.Comment{
		AuthorId:  comment.UserID,
		ParentId:  comment.ParentID,
		CreatedAt: time.Now().UTC(),
		EditedAt:  time.Now().UTC(),
		Content:   comment.Content,
	}
	dComment.PostId = &comment.PostID
	if comment.ParentID != nil {
		dComment.PostId = nil
	}
	return dComment
}

func DomainToDtoComment(comment domain.Comment) *dto.CommentDto {
	return &dto.CommentDto{
		Id:        comment.Id,
		ParentId:  comment.ParentId,
		AuthorId:  comment.AuthorId,
		PostId:    comment.PostId,
		CreatedAt: comment.CreatedAt,
		EditedAt:  comment.EditedAt,
		Content:   comment.Content,
	}
}

func DtoToGqlComment(comment dto.CommentDto) *model.Comment {
	return &model.Comment{
		ID:        comment.Id,
		ParentID:  comment.ParentId,
		PostID:    comment.PostId,
		CreatedAt: &comment.CreatedAt,
		EditedAt:  &comment.EditedAt,
		Content:   &comment.Content,
	}
}

func GqlToDtoCommentInput(commentInput model.CommentCreateInput) *dto.CommentCreateInput {
	return &dto.CommentCreateInput{
		PostID:   commentInput.PostID,
		ParentID: commentInput.ParentID,
		UserID:   commentInput.UserID,
		Content:  commentInput.Content,
	}
}

func GqlToDtoPostInput(postInput model.PostCreateInput) *dto.PostCreateDto {
	return &dto.PostCreateDto{
		UserId:          postInput.UserID,
		Title:           postInput.Title,
		Content:         postInput.Content,
		CommentsAllowed: postInput.CommentsAllowed,
	}
}

func GqlToDtoPostUpdateInput(postInput model.PostUpdateInput) *dto.PostUpdateDto {
	return &dto.PostUpdateDto{
		PostId:          postInput.PostID,
		UserId:          postInput.UserID,
		Title:           postInput.Title,
		Content:         postInput.Content,
		CommentsAllowed: postInput.CommentsAllowed,
	}
}

func GqlToDtoCommentUpdateInput(commentInput model.CommentUpdateInput) *dto.CommentUpdateInput {
	return &dto.CommentUpdateInput{
		UserID:    commentInput.UserID,
		CommentID: commentInput.CommentID,
		Content:   commentInput.Content,
	}
}

func DomainCommentsToGqlEdges(comments []domain.Comment) []*model.CommentEdge {
	edges := make([]*model.CommentEdge, len(comments))

	for i, comment := range comments {
		cursor := comment.Id.String()
		edges[i] = &model.CommentEdge{
			Comment: DomainToGqlComment(comment),
			Cursor:  &cursor,
		}
	}

	return edges
}

func DtoPostsToGqlEdges(posts []dto.PostDto) []*model.PostEdge {
	edges := make([]*model.PostEdge, len(posts))

	for i, post := range posts {
		cursor := post.Id.String()
		edges[i] = &model.PostEdge{
			Post:   DtoToGqlPost(post),
			Cursor: &cursor,
		}
	}

	return edges
}

func DomainCommentsToGqlPageInfo(comments []domain.Comment) *model.PageInfo {
	hasNextPage := false
	if len(comments) == 0 {
		return &model.PageInfo{
			HasNextPage: &hasNextPage,
			EndCursor:   nil,
		}
	}
	lastId := comments[len(comments)-1].Id.String()
	hasNextPage = true
	return &model.PageInfo{
		HasNextPage: &hasNextPage,
		EndCursor:   &lastId,
	}
}

func DtoPostsToGqlPageInfo(posts []dto.PostDto) *model.PageInfo {
	hasNextPage := false
	if len(posts) == 0 {
		return &model.PageInfo{
			HasNextPage: &hasNextPage,
			EndCursor:   nil,
		}
	}
	lastId := posts[len(posts)-1].Id.String()
	hasNextPage = true

	return &model.PageInfo{
		HasNextPage: &hasNextPage,
		EndCursor:   &lastId,
	}
}
