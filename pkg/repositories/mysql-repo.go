package repositories

import (
	"github.com/io-m/hexagonal/pkg/entities"
	"github.com/io-m/hexagonal/pkg/utils/error"
)

type repository struct{}

// NewMySQLRepository is constructor function for making new instances of Repository interface
func NewMySQLRepository() Repository {
	return &repository{}
}

var allPosts []*entities.Post

// Save creates new post inside MySQL DB
func (*repository) Save(post *entities.Post) (*entities.Post, *error.Response) {
	incommingPost := entities.Post{
		ID:    post.ID,
		Title: post.Title,
		Body:  post.Body,
	}
	allPosts = append(allPosts, &incommingPost)
	return &incommingPost, nil
}

// GetAll fetches all data from MySQL DB
func (*repository) GetAll() ([]*entities.Post, *error.Response) {
	return allPosts, nil
}
