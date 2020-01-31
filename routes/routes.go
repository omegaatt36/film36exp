package routes

import (
	"film36exp/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Method     string
	Pattern    string
	Handler    http.HandlerFunc
	Middleware mux.MiddlewareFunc
}

var routes []Route

func init() {
	register("POST", "/event", controllers.CreateEvent, nil)
	register("GET", "/events", controllers.GetAllEvents, nil)
	register("GET", "/events/{id}", controllers.GetOneEvent, nil)
	register("PATCH", "/events/{id}", controllers.UpdateEvent, nil)
	register("DELETE", "/events/{id}", controllers.DeleteEvent, nil)
}

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
	routes = append(routes, Route{method, pattern, handler, middleware})
}
