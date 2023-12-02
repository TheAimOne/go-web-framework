package main

import (
	"log"
	"net/http"

	"github.com/TheAimOne/go-web-framework/endpoint"
	"github.com/TheAimOne/go-web-framework/handler"
	"github.com/TheAimOne/go-web-framework/server"
)

func main() {
	s := server.NewServer()

	s.AddHandler(endpoint.Endpoint{
		Path:    "/test",
		Method:  "POST",
		Handler: handler.CreateTestHandler,
	})

	s.Start()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
