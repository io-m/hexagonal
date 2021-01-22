package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/io-m/hexagonal/pkg/entities"
	"github.com/io-m/hexagonal/pkg/services"
	"github.com/io-m/hexagonal/pkg/utils/error"
)

// PostHandler interface defines method for handling certain endpoints
// Endpoints are defined in rputer package
type PostHandler interface {
	GetAllPosts(w http.ResponseWriter, r *http.Request)
	CreatePost(w http.ResponseWriter, r *http.Request)
}

var postService = services.NewPostService()

// GetAllPosts is handler for GET http method
func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	allPosts, err := postService.FetchAllService()
	if err != nil {
		w.WriteHeader(err.StatusCode)
		json.NewEncoder(w).Encode(err)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(allPosts)
}

// CreatePost is handler for POST http method
func CreatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	newPost := entities.Post{}
	if err := json.NewDecoder(r.Body).Decode(&newPost); err != nil {
		theErr := error.Back(http.StatusBadRequest, "Bad request")
		w.WriteHeader(theErr.StatusCode)
		json.NewEncoder(w).Encode(theErr)
	}
	if err := postService.Validate(&newPost); err != nil {
		w.WriteHeader(err.StatusCode)
		json.NewEncoder(w).Encode(err.Msg)
		return
	}
	createdPost, theErr := postService.CreateService(&newPost)
	if theErr != nil {
		w.WriteHeader(theErr.StatusCode)
		json.NewEncoder(w).Encode(theErr)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(createdPost)
}
