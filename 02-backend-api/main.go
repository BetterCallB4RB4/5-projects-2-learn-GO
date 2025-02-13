package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
)

var logger *slog.Logger

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	logger.Info("init server")

	router := http.NewServeMux()
	router.HandleFunc("/", welcome)
	err := http.ListenAndServe("localhost:8080", router)
	if err != nil {
		log.Fatal("internal server error")
	}
}

func welcome(w http.ResponseWriter, r *http.Request) {
	logger.Info("received request at home endpoint", "method", r.Method, "url", r.URL)
	fmt.Fprintf(w, "Hello, World!\n")
}
