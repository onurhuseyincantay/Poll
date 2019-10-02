package api

import (
	"api/config"
	"api/router"
	"auto"
	"fmt"
	"log"
	"net/http"
)

func Run() {
	config.Load()

	auto.Load()
	fmt.Println("Server Running on :", config.PORT)
	listen(config.PORT)
}

func listen(port int) {
	r := router.NEW()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}
