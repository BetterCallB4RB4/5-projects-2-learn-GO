package main

import (
	"encoding/json"
	"net/http"
)

type Operation struct {
	Number1 int `json:"number1"`
	Number2 int `json:"number2"`
}

type ResultResponse struct {
	Result int `json:"result"`
}

func add(w http.ResponseWriter, r *http.Request) {
	numbers := formatRequest(r)
	result := numbers.Number1 + numbers.Number2
	w.Header().Set("Content-Type", "application/json")
	w.Write(formatResponse(result))
}

func subtract(w http.ResponseWriter, r *http.Request) {
	numbers := formatRequest(r)
	result := numbers.Number1 - numbers.Number2
	w.Header().Set("Content-Type", "application/json")
	w.Write(formatResponse(result))
}

func multiply(w http.ResponseWriter, r *http.Request) {
	numbers := formatRequest(r)
	result := numbers.Number1 * numbers.Number2
	w.Header().Set("Content-Type", "application/json")
	w.Write(formatResponse(result))
}

func divide(w http.ResponseWriter, r *http.Request) {
	numbers := formatRequest(r)
	result := numbers.Number1 / numbers.Number2
	w.Header().Set("Content-Type", "application/json")
	w.Write(formatResponse(result))
}

func module(w http.ResponseWriter, r *http.Request) {
	numbers := formatRequest(r)
	result := numbers.Number1 % numbers.Number2
	w.Header().Set("Content-Type", "application/json")
	w.Write(formatResponse(result))
}

func formatRequest(r *http.Request) Operation {
	var numbers Operation
	err := json.NewDecoder(r.Body).Decode(&numbers)
	if err != nil {
		logger.Error("couldn't read the request body, cannot perform any operation")
	}
	return numbers
}

func formatResponse(result int) []byte {
	response := ResultResponse{Result: result}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		logger.Error("couldn't marshal the response to json")
	}
	return append(jsonResponse, '\n')
}
