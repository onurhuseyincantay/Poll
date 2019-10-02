package routes

import (
	"fmt"
	"net/http"
	"api/middlewares"
	"github.com/gorilla/mux"
)

type Route struct {
	URI        string
	MethodType string
	Handler    func(http.ResponseWriter, *http.Request)
}

func Load() []Route {
	routes := usersRoutes
	return routes
}

func SetupRoutes(r *mux.Router) *mux.Router {
	for _, route := range Load() {
		r.HandleFunc(route.URI, route.Handler).Methods(route.MethodType)
		fmt.Println(route.URI)
	}
	return r
}

func SetupRoutesWithMiddleWare(r *mux.Router) *mux.Router {
	for _, route := range Load() {
		r.HandleFunc(route.URI,
			middlewares.SetMiddlewareLogger(
				middlewares.SetMiddlewareJSON(route.Handler)),
		).Methods(route.MethodType)
		fmt.Println(route.URI)
	}
	return r
}
