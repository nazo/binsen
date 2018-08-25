package repositories

import (
	"context"
	"database/sql"

	"github.com/nazo/binsen/server/app/orm"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

// PostsRepository todo
type PostsRepository interface {
	GetPost(id int64) (*orm.Post, error)
	GetPosts(workspace *orm.Workspace, offset int, limit int) ([]*orm.Post, error)
	CreatePost(workspace *orm.Workspace, title string, body string, user *orm.User) (*orm.Post, error)
	UpdatePost(post *orm.Post, title string, body string, user *orm.User) (*orm.Post, error)
}

// postsRepository todo
type postsRepository struct {
	PostsRepository
	db *sql.DB
}

// NewPostsRepository todo
func NewPostsRepository(db *sql.DB) PostsRepository {
	return &postsRepository{db: db}
}

// GetPosts get posts
func (r *postsRepository) GetPosts(workspace *orm.Workspace, offset int, limit int) ([]*orm.Post, error) {
	return orm.Posts(qm.Where("workspace_id = ?", workspace.ID), qm.OrderBy("updated_at DESC"), qm.Offset(offset), qm.Limit(limit)).All(context.Background(), r.db)
}

// GetPost get post
func (r *postsRepository) GetPost(id int64) (*orm.Post, error) {
	return orm.Posts(qm.Where("id = ?", id)).One(context.Background(), r.db)
}

// CreatePost create post
func (r *postsRepository) CreatePost(workspace *orm.Workspace, title string, body string, user *orm.User) (*orm.Post, error) {
	post := &orm.Post{
		WorkspaceID:   workspace.ID,
		Title:         title,
		Body:          body,
		CreatorUserID: user.ID,
		EditorUserID:  user.ID,
		EditStatus:    1,
	}
	err := post.Insert(context.Background(), r.db, boil.Infer())
	return post, err
}

// UpdatePost update post
func (r *postsRepository) UpdatePost(post *orm.Post, title string, body string, user *orm.User) (*orm.Post, error) {
	post.Title = title
	post.Body = body
	post.EditorUserID = user.ID
	_, err := post.Update(context.Background(), r.db, boil.Infer())
	return post, err
}
