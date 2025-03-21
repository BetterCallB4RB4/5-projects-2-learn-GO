package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Operation struct {
	Number1 int `json:"number1"`
	Number2 int `json:"number2"`
}

func add(w http.ResponseWriter, r *http.Request) {
	var numbers Operation
	err := json.NewDecoder(r.Body).Decode(&numbers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		logger.Error("couldn't read the request body, cannot perform any operation")
	}
	result := numbers.Number1 + numbers.Number2
	w.Write([]byte(strconv.Itoa(result)))
}

func subtract(w http.ResponseWriter, r *http.Request) {
	var numbers Operation
	err := json.NewDecoder(r.Body).Decode(&numbers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		logger.Error("couldn't read the request body, cannot perform any operation")
	}
	result := numbers.Number1 - numbers.Number2
	w.Write([]byte(strconv.Itoa(result)))
}

func multiply(w http.ResponseWriter, r *http.Request) {
	var numbers Operation
	err := json.NewDecoder(r.Body).Decode(&numbers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		logger.Error("couldn't read the request body, cannot perform any operation")
	}
	result := numbers.Number1 * numbers.Number2
	w.Write([]byte(strconv.Itoa(result)))
}

func divide(w http.ResponseWriter, r *http.Request) {
	var numbers Operation
	err := json.NewDecoder(r.Body).Decode(&numbers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		logger.Error("couldn't read the request body, cannot perform any operation")
	}
	result := numbers.Number1 / numbers.Number2
	w.Write([]byte(strconv.Itoa(result)))
}
