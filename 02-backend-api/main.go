package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
)

var logger *slog.Logger

type Operation struct {
	Number1 int `json:"number1"`
	Number2 int `json:"number2"`
}

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	logger.Info("init server")

	router := http.NewServeMux()
	err := http.ListenAndServe("localhost:8080", router)
	if err != nil {
		log.Fatal("internal server error")
	}
}

func welcome(w http.ResponseWriter, r *http.Request) {
	logger.Info("received request at home endpoint", "method", r.Method, "url", r.URL)
	fmt.Fprintf(w, "Hello, World!\n")
}

func addOperation(operand Operation) int {
	return operand.Number1 + operand.Number2
}

func subtractionOperation(operand Operation) int {
	return operand.Number1 - operand.Number2
}

func multiplicationOperation(operand Operation) int {
	return operand.Number1 * operand.Number2
}

func divisionOperation(operand Operation) int {
	return operand.Number1 / operand.Number2
}
