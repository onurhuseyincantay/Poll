package routes

import (
	"api/controllers"
	"net/http"
)

var pollRoutes = []Route{
	Route{
		URI:        "/api/polls",
		MethodType: http.MethodGet,
		Handler:    controllers.GetPolls,
	},
	Route{
		URI:        "/api/polls/{id}",
		MethodType: http.MethodGet,
		Handler:    controllers.GetPollByID,
	},
	Route{
		URI:        "/api/polls",
		Queries:    "userId",
		MethodType: http.MethodGet,
		Handler:    controllers.GetPollsByUserId,
	},
}