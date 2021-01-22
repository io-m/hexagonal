package services

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/io-m/hexagonal/pkg/entities"
	"github.com/io-m/hexagonal/pkg/repositories"
	"github.com/io-m/hexagonal/pkg/utils/error"
)

// PostService interface defines methods for
// validating incomming posts
// creating new posts received from handlers
// fetching all posts by asking repository to
// interact with DB
type PostService interface {
	Validate(post *entities.Post) *error.Response
	CreateService(post *entities.Post) (*entities.Post, *error.Response)
	FetchAllService() ([]*entities.Post, *error.Response)
}

type service struct{}

// NewPostService is constructor function for creating of new instances
// of PostService interface
func NewPostService() PostService {
	return &service{}
}

func (*service) Validate(post *entities.Post) *error.Response {
	if post == nil {
		return error.Back(http.StatusBadRequest, "Post cannot be empty")
	}

	if post.Title == "" {
		return error.Back(http.StatusBadRequest, "Post title cannot be empty")
	}

	if post.Body == "" {
		return error.Back(http.StatusBadRequest, "Post body cannot be empty")
	}

	return nil
}

var repo = repositories.NewMySQLRepository()

func (*service) CreateService(post *entities.Post) (*entities.Post, *error.Response) {
	post.ID = uuid.New()
	return repo.Save(post)
}

func (*service) FetchAllService() ([]*entities.Post, *error.Response) {
	return repo.GetAll()
}
