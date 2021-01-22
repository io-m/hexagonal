package endpoints

import (
	"github.com/io-m/hexagonal/pkg/handlers"
	"github.com/io-m/hexagonal/pkg/routers"
)

// RunApp is an entry point to the app
// It calls endpoints and running server
// It is called from cmd/main.go
func RunApp() {
	router := routers.NewMuxRouter()
	router.Get("/api/v1/post", handlers.GetAllPosts)
	router.Post("/api/v1/post", handlers.CreatePost)
	router.ListenAndServe(":6600")
}
