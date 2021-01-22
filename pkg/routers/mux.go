package routers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type gorillaMux struct{}

// NewMuxRouter is constructor function for creating new instances of Gorilla mux router
func NewMuxRouter() Router {
	return &gorillaMux{}
}

var muxRouter = mux.NewRouter()

func (*gorillaMux) Get(path string, handler func(w http.ResponseWriter, r *http.Request)) {
	muxRouter.HandleFunc(path, handler).Methods(http.MethodGet)
}

func (*gorillaMux) Post(path string, handler func(w http.ResponseWriter, r *http.Request)) {
	muxRouter.HandleFunc(path, handler).Methods(http.MethodPost)
}

func (*gorillaMux) ListenAndServe(port string) {
	log.Fatal(http.ListenAndServe(port, muxRouter))
	log.Printf("Server is up and running on PORT : %v", port)
}
