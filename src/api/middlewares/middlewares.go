package middlewares

import (
	"log"
	"net/http"
)

func SetMiddlewareLogger(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		log.Printf("%s %s%s %s", request.Method, request.Host, request.RequestURI, request.Proto)
		next(writer, request)
	}
}

func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type","application/json")
		next(writer,request)
	}
}
