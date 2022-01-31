package main

import (
	"log"
	"net"
	"net/http"
	"time"
	"todolist/internal/user"

	"github.com/julienschmidt/httprouter"
)

func main() {
	log.Println("create router")
	router := httprouter.New()
	log.Println("register list router")

	handler := user.NewHandler()
	handler.Register(router)
	start(router)

}

func start(router *httprouter.Router) {
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(server.Serve(listener))
}
