package repositories

import (
	"github.com/io-m/hexagonal/pkg/entities"
	"github.com/io-m/hexagonal/pkg/utils/error"
)

// Repository interface defines functions for interacting with DB
type Repository interface {
	Save(post *entities.Post) (*entities.Post, *error.Response)
	GetAll() ([]*entities.Post, *error.Response)
}
