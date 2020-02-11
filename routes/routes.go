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
	register("GET", "/", controllers.GetDefault, nil)
	register("POST", "/Createfilm", controllers.CreateOneFilm, nil)
	register("GET", "/films", controllers.GetAllFilms, nil)
	register("GET", "/films/{filmID}", controllers.GetOneFilm, nil)
	register("PATCH", "/films/{filmID}", controllers.UpdateFilm, nil)
	register("DELETE", "/films/{filmID}", controllers.DeleteFilm, nil)

	register("POST", "/CreatePic/{filmID}", controllers.CreatePic, nil)
	register("PATCH", "/pics/{filmID}/{picID}", controllers.UpdatePic, nil)
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
