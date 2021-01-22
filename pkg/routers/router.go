package routers

import "net/http"

// Router interface defines http methods, endpoints and handlers
type Router interface {
	Get(path string, handler func(w http.ResponseWriter, r *http.Request))
	Post(path string, handler func(w http.ResponseWriter, r *http.Request))
	ListenAndServe(port string)
}
