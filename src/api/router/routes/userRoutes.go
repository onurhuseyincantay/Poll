package routes

import (
	"api/controllers"
	"net/http"
)

var usersRoutes = []Route{
	Route{
		URI:        "/api/user",
		MethodType: http.MethodGet,
		Handler:    controllers.GetUsers,
	},
	Route{
		URI:        "/api/user/{id}",
		MethodType: http.MethodGet,
		Handler:    controllers.GetUserById,
	},
	Route{
		URI:        "/api/user/register",
		MethodType: http.MethodPost,
		Handler:    controllers.RegisterUser,
	},
	Route{
		URI:        "/api/user/login",
		MethodType: http.MethodPost,
		Handler:    controllers.LoginUser,
	},
	Route{
		URI:        "/api/user/deleteAccount",
		MethodType: http.MethodPost,
		Handler:    controllers.DeleteUser,
	},
	Route{
		URI:        "/api/user/updateProfile/{id}",
		MethodType: http.MethodPut,
		Handler:    controllers.UpdateProfile,
	},
}
