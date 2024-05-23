package inmem

import (
	"fmt"
	"sync"

	"github.com/PonomarevAlexxander/graphql-forum/internal/domain"
	"github.com/PonomarevAlexxander/graphql-forum/internal/errlib"
	"github.com/google/uuid"
)

type db struct {
	uM                 sync.RWMutex
	pM                 sync.RWMutex
	cM                 sync.RWMutex
	users              map[uuid.UUID]domain.User
	posts              map[uuid.UUID]domain.Post
	comments           map[uuid.UUID]domain.Comment
	commentsToComments map[uuid.UUID][]uuid.UUID
	postsToComments    map[uuid.UUID][]uuid.UUID
	commentToUsers     map[uuid.UUID]uuid.UUID
	postsToUsers       map[uuid.UUID]uuid.UUID
}

func NewDb() *db {
	return &db{
		users:              make(map[uuid.UUID]domain.User),
		posts:              make(map[uuid.UUID]domain.Post),
		comments:           make(map[uuid.UUID]domain.Comment),
		commentsToComments: make(map[uuid.UUID][]uuid.UUID),
		postsToComments:    make(map[uuid.UUID][]uuid.UUID),
		commentToUsers:     make(map[uuid.UUID]uuid.UUID),
		postsToUsers:       make(map[uuid.UUID]uuid.UUID),
	}
}

func (d *db) NewUser(user domain.User) error {
	d.uM.Lock()
	defer d.uM.Unlock()

	_, exists := d.users[user.Id]
	if exists {
		return fmt.Errorf("%w, user with this id already exists", errlib.ErrResourceAlreadyExists)
	}

	d.users[user.Id] = user
	return nil
}

func (d *db) GetUser(id uuid.UUID) (domain.User, error) {
	d.uM.RLock()
	defer d.uM.RUnlock()

	user, ok := d.users[id]
	if !ok {
		return domain.User{}, fmt.Errorf("%w, user not found", errlib.ErrNotFound)
	}

	return user, nil
}

func (d *db) GetUserByCommentId(commentId uuid.UUID) (domain.User, error) {
	d.cM.RLock()
	userId, ok := d.commentToUsers[commentId]
	if !ok {
		return domain.User{}, fmt.Errorf("%w, comment not found", errlib.ErrNotFound)
	}
	d.cM.RUnlock()

	d.uM.RLock()
	defer d.uM.RUnlock()
	user, ok := d.users[userId]
	if !ok {
		return domain.User{}, fmt.Errorf("%w, user not found", errlib.ErrNotFound)
	}

	return user, nil
}

func (d *db) GetUserByPostId(postId uuid.UUID) (domain.User, error) {
	d.pM.RLock()
	userId, ok := d.postsToUsers[postId]
	if !ok {
		return domain.User{}, fmt.Errorf("%w, post not found", errlib.ErrNotFound)
	}
	d.pM.RUnlock()

	d.uM.RLock()
	defer d.uM.RUnlock()
	user, ok := d.users[userId]
	if !ok {
		return domain.User{}, fmt.Errorf("%w, user not found", errlib.ErrNotFound)
	}

	return user, nil
}

func (d *db) NewPost(post domain.Post) error {
	d.pM.Lock()
	defer d.pM.Unlock()

	_, exists := d.posts[post.Id]
	if exists {
		return fmt.Errorf("%w, post with this id already exists", errlib.ErrResourceAlreadyExists)
	}

	d.posts[post.Id] = post
	d.postsToUsers[post.Id] = post.AuthorId
	return nil
}

func (d *db) UpdatePost(post domain.Post) error {
	d.pM.Lock()
	defer d.pM.Unlock()

	_, exists := d.posts[post.Id]
	if !exists {
		return fmt.Errorf("%w, post with this id doesn't exist", errlib.ErrResourceAlreadyExists)
	}

	d.posts[post.Id] = post

	return nil
}

func (d *db) GetPost(id uuid.UUID) (domain.Post, error) {
	d.pM.RLock()
	defer d.pM.RUnlock()

	post, ok := d.posts[id]
	if !ok {
		return domain.Post{}, fmt.Errorf("%w, post not found", errlib.ErrNotFound)
	}

	return post, nil
}

func (d *db) CountPosts() uint {
	d.pM.RLock()
	defer d.pM.RUnlock()

	return uint(len(d.posts))
}

func (d *db) GetPosts() ([]domain.Post, error) {
	d.pM.RLock()
	defer d.pM.RUnlock()

	posts := make([]domain.Post, len(d.posts))
	index := 0
	for _, post := range d.posts {
		posts[index] = post
		index++
	}

	return posts, nil
}

func (d *db) GetComments() ([]domain.Comment, error) {
	d.cM.RLock()
	defer d.cM.RUnlock()

	comments := make([]domain.Comment, len(d.comments))
	index := 0
	for _, comment := range d.comments {
		comments[index] = comment
		index++
	}

	return comments, nil
}

func (d *db) NewComment(comment domain.Comment) error {
	d.cM.Lock()
	defer d.cM.Unlock()

	_, ok := d.comments[comment.Id]
	if ok {
		return fmt.Errorf("%w, comment with this id already exists", errlib.ErrResourceAlreadyExists)
	}

	d.comments[comment.Id] = comment

	d.commentToUsers[comment.Id] = comment.AuthorId
	if comment.PostId != nil {
		d.postsToComments[*comment.PostId] = append(d.postsToComments[*comment.PostId], comment.Id)
	} else if comment.ParentId != nil {
		d.commentsToComments[*comment.ParentId] = append(d.commentsToComments[*comment.ParentId], comment.Id)
	} else {
		return fmt.Errorf("%w, comment should have postId or parentId", errlib.ErrBadRequest)
	}

	return nil
}

func (d *db) UpdateComment(comment domain.Comment) error {
	d.cM.Lock()
	defer d.cM.Unlock()

	_, ok := d.comments[comment.Id]
	if !ok {
		return fmt.Errorf("%w, comment with this id doesn't exist", errlib.ErrResourceAlreadyExists)
	}

	d.comments[comment.Id] = comment

	return nil
}

func (d *db) GetComment(id uuid.UUID) (domain.Comment, error) {
	d.cM.RLock()
	defer d.cM.RUnlock()

	comment, ok := d.comments[id]
	if !ok {
		return domain.Comment{}, fmt.Errorf("%w, comment not found", errlib.ErrNotFound)
	}

	return comment, nil
}

func (d *db) CountCommentsByPostId(postId uuid.UUID) (uint, error) {
	d.cM.RLock()
	defer d.cM.RUnlock()

	comments, ok := d.postsToComments[postId]
	if !ok {
		return 0, fmt.Errorf("%w, post not found", errlib.ErrNotFound)
	}

	return uint(len(comments)), nil
}

func (d *db) CountCommentsByCommentId(commentId uuid.UUID) (uint, error) {
	d.cM.RLock()
	defer d.cM.RUnlock()

	comments, ok := d.commentsToComments[commentId]
	if !ok {
		return 0, fmt.Errorf("%w, comment not found", errlib.ErrNotFound)
	}

	return uint(len(comments)), nil
}

func (d *db) GetPostComments(postId uuid.UUID) ([]domain.Comment, error) {
	_, err := d.GetPost(postId)
	if err != nil {
		return nil, err
	}

	d.cM.RLock()
	defer d.cM.RUnlock()

	commentIds, ok := d.postsToComments[postId]
	if !ok || len(commentIds) == 0 {
		return []domain.Comment{}, nil
	}

	comments := make([]domain.Comment, len(commentIds))

	for i, id := range commentIds {
		comment, ok := d.comments[id]
		if !ok {
			return nil, fmt.Errorf("%w, comment not found", errlib.ErrNotFound)
		}

		comments[i] = comment
	}

	return comments, nil
}

func (d *db) GetCommentComments(commentId uuid.UUID) ([]domain.Comment, error) {
	_, err := d.GetComment(commentId)
	if err != nil {
		return nil, err
	}

	d.cM.RLock()
	defer d.cM.RUnlock()

	commentIds, ok := d.commentsToComments[commentId]
	if !ok || len(commentIds) == 0 {
		return []domain.Comment{}, nil
	}

	comments := make([]domain.Comment, len(commentIds))

	for i, id := range commentIds {
		comment, ok := d.comments[id]
		if !ok {
			return nil, fmt.Errorf("%w, comment not found", errlib.ErrNotFound)
		}

		comments[i] = comment
	}

	return comments, nil
}
