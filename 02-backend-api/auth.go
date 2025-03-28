package main

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"time"
)

// in a prod env you have to store this information in a secure way
var tokens = make(map[string]time.Time)

// TODO: forse sarebbe carine usare direttamente qua una struttura invece che fare una conversione solo durante la risposta

type tokenResponse struct {
	Token  string `json:"token"`
	Expiry string `json:"expiry"`
}

func generateToken() (string, error) {
	token := make([]byte, 32)  // adjust lenght as needed
	_, err := rand.Read(token) // this read some bytes from a source that is "cryptographically" secure random number generator
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(token), nil
}

func authenticate(token string) bool {
	expiry, tokenFound := tokens[token]
	if !tokenFound {
		return false
	} else if time.Now().After(expiry) {
		delete(tokens, token)
		return false
	}
	return true
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	// In a real app, validate username/password
	token, err := generateToken()
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	expiry := time.Now().Add(1 * time.Minute)
	tokens[token] = expiry
	logger.Info("release token: " + token)

	response := tokenResponse{
		Token:  token,
		Expiry: expiry.Format(time.RFC3339),
	}

	// encode the json response fot the client
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

// curl -H "Content-Type: application/json" -H "Authorization: $(curl -H "Content-Type: application/json" -d '{"number1": 10, "number2": 5}' http://localhost:8080/login | jq ".token" -r)" -d '{"number1": 10, "number2": 5}' http://localhost:8080/add
