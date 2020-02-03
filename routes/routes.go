package routes

import (
	"film36exp/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

type route struct {
	Method     string
	Pattern    string
	Handler    http.HandlerFunc
	Middleware mux.MiddlewareFunc
}

var routes []route

func init() {
	register("POST", "/Createfilm", controllers.CreateOneFilm, nil)
	register("GET", "/films", controllers.GetAllFilms, nil)
	register("GET", "/films/{id}", controllers.GetOneFilm, nil)
	register("PATCH", "/films/{id}", controllers.UpdateFilm, nil)
	register("DELETE", "/films/{id}", controllers.DeleteFilm, nil)
}

// NewRouter create one 'mux.NewRouter()' and register handle funcs.
func NewRouter() *mux.Router {
	r := mux.NewRouter()
	for _, route := range routes {
		r.HandleFunc(route.Pattern, route.Handler).Methods(route.Method)

		if route.Middleware != nil {
			r.Use(route.Middleware)
		}
	}
	return r
}

func register(method, pattern string, handler http.HandlerFunc, middleware mux.MiddlewareFunc) {
	routes = append(routes, route{method, pattern, handler, middleware})
}
