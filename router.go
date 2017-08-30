package pot

import "github.com/gorilla/mux"

type Router struct {
	mux.Router
}

func NewRouter() *Router {
	return &Router{
		Router: *mux.NewRouter(),
	}
}
