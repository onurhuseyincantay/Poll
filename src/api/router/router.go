package router

import (
	"api/router/routes"

	"github.com/gorilla/mux"
)

func NEW() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	return routes.SetupRoutesWithMiddleWare(r)
}
