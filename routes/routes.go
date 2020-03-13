package routes

import (
	"film36exp/auth"
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
	register("POST", "/Createfilm", controllers.CreateOneFilm, auth.TokenMiddleware)
	register("GET", "/films", controllers.GetAllFilms, auth.TokenMiddleware)
	register("GET", "/films/{filmID}", controllers.GetOneFilm, auth.TokenMiddleware)
	register("PATCH", "/films/{filmID}", controllers.UpdateFilm, auth.TokenMiddleware)
	register("DELETE", "/films/{filmID}", controllers.DeleteFilm, auth.TokenMiddleware)

	register("POST", "/CreatePic/{filmID}", controllers.CreatePic, auth.TokenMiddleware)
	register("GET", "/pics/{filmID}", controllers.GetPics, auth.TokenMiddleware)
	register("PATCH", "/pics/{picID}", controllers.UpdatePic, auth.TokenMiddleware)

	register("POST", "/user/register", controllers.Register, nil)
	register("POST", "/user/login", controllers.Login, nil)
}

// NewRouter create one 'mux.NewRouter()' and register handle funcs.
func NewRouter() *mux.Router {
	r := mux.NewRouter()
	for _, route := range routes {
		if route.Middleware == nil {
			r.HandleFunc(route.Pattern, route.Handler).Methods(route.Method)
		} else {
			r.Handle(route.Pattern, route.Middleware(route.Handler)).Methods(route.Method)
		}
	}
	return r
}

func register(method, pattern string, handler http.HandlerFunc, middleware mux.MiddlewareFunc) {
	routes = append(routes, route{method, pattern, handler, middleware})
}
