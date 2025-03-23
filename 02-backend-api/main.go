package main

import (
	"log/slog"
	"net/http"
	"os"
)

var logger *slog.Logger

func main() {
	logger = slog.New(slog.NewTextHandler(os.Stdout, nil))
	logger.Info("init server")

	router := http.NewServeMux()

	router.Handle("/add", applyMiddleware(add))
	router.Handle("/subtract", applyMiddleware(subtract))
	router.Handle("/multiply", applyMiddleware(multiply))
	router.Handle("/divide", applyMiddleware(divide))

	// #### start serving the request ####
	err := http.ListenAndServe("localhost:8080", router) // questo e' il server http a cui passo il router per dire come gestire le chimate
	if err != nil {
		logger.Error("internal server error on starting the multiplexer")
	}
}
