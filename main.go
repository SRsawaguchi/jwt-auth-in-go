package main

import (
	"context"
	"log"
	"net/http"

	"github.com/SRsawaguchi/jwt-auth-in-go/internal/server"
)

func main() {
	port := "8080"
	server := server.Server{}
	if err := server.Init(context.Background()); err != nil {
		log.Fatalf(err.Error())
	}

	log.Printf("The server starting at http://localhost:%s/", port)
	log.Fatal(http.ListenAndServe(":"+port, server.Handler()))
}
