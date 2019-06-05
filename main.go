package main

import (
	"net/http"

	"github.com/ddo/go-vue-handler"
)

const (
	port = "8080"
	publicDir = "./public"
)

func main() {
	server := &http.Server{
		Addr:    ":" + port,
		Handler: vue.Handler(publicDir),
	}
	err := server.ListenAndServe()
	panic(err)
}
