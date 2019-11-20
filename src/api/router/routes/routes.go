package routes

import (
	"fmt"
	"net/http"
	"api/middlewares"
	"github.com/gorilla/mux"
)

type Route struct {
	URI        string
	Queries    string
	MethodType string
	Handler    func(http.ResponseWriter, *http.Request)
}

func Load() []Route {
	routes := append(usersRoutes, pollRoutes...)
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
		if  route.Queries != "" {
			query := fmt.Sprintf("{%s}", route.Queries)
			r.HandleFunc(route.URI,
				middlewares.SetMiddlewareLogger(
					middlewares.SetMiddlewareJSON(route.Handler)),
			).Methods(route.MethodType).Queries(route.Queries,query)
		} else {
			r.HandleFunc(route.URI,
				middlewares.SetMiddlewareLogger(
					middlewares.SetMiddlewareJSON(route.Handler)),
			).Methods(route.MethodType)

		}
		fmt.Println(route.URI)
	}
	return r
}
