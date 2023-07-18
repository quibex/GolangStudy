package main

import (
	"RestAPI/internal/user"
	"RestAPI/pkg/logging"
	"github.com/julienschmidt/httprouter"
	"log"
	"net"
	"net/http"
	"time"
)

func main() {
	logger := logging.GetLogger()
	logger.Info("create router")
	router := httprouter.New()

	logger.Info("register user handler")
	handler := user.NewHandler(logger)
	handler.Register(router)
	start(router)

}

func start(router *httprouter.Router) {
	log.Println("start app")

	listener, err := net.Listen("tcp", ":3000")
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Handler:           router,
		WriteTimeout:      15 * time.Second,
		ReadHeaderTimeout: 15 * time.Second,
	}

	log.Println("server is listening port localhost:3000")
	log.Fatalln(server.Serve(listener))
}
