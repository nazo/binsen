package services

import (
	"database/sql"

	"github.com/nazo/binsen/server/app/orm"
	"github.com/nazo/binsen/server/app/repositories"
)

// PostService todo
type PostService interface {
	CreatePost(workspace *orm.Workspace, title string, body string, user *orm.User) (*orm.Post, error)
	UpdatePost(post *orm.Post, title string, body string, user *orm.User) (*orm.Post, error)
	GetPosts(workspace *orm.Workspace, page int) ([]*orm.Post, error)
	GetPost(id int64) (*orm.Post, error)
}

// postService todo
type postService struct {
	PostService
	postsRepository repositories.PostsRepository
}

// NewPostService todo
func NewPostService(db *sql.DB) PostService {
	return &postService{
		postsRepository: repositories.NewPostsRepository(db),
	}
}

// CreatePost create post
func (s *postService) CreatePost(workspace *orm.Workspace, title string, body string, user *orm.User) (*orm.Post, error) {
	return s.postsRepository.CreatePost(workspace, title, body, user)
}

// UpdatePost update post
func (s *postService) UpdatePost(post *orm.Post, title string, body string, user *orm.User) (*orm.Post, error) {
	return s.postsRepository.UpdatePost(post, title, body, user)
}

// GetPosts get posts
func (s *postService) GetPosts(workspace *orm.Workspace, page int) ([]*orm.Post, error) {
	limit := 20
	offset := (page - 1) * limit
	posts, err := s.postsRepository.GetPosts(workspace, offset, limit)
	if err == sql.ErrNoRows {
		return []*orm.Post{}, nil
	} else if err != nil {
		return nil, err
	}
	return posts, nil
}

// GetPost get post
func (s *postService) GetPost(id int64) (*orm.Post, error) {
	return s.postsRepository.GetPost(id)
}
