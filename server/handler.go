package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/rmh-softengineer/locqube/api/model"
)

// In-memory storage for tokens
var userTokens = map[string]string{}

func handleLoginToFacebook(w http.ResponseWriter, r *http.Request) {
	var req model.LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid login request", http.StatusBadRequest)
		return
	}

	token, err := env.FacebookService.Login(req.AccessToken)
	if err != nil {
		http.Error(w, "invalid Facebook token", http.StatusUnauthorized)
		return
	}

	userTokens[*token] = req.AccessToken // Save the token in a more persistent and secure way

	response := model.LoginResponse{Token: *token}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func handlePostToFacebook(w http.ResponseWriter, r *http.Request) {
	token, err := extractToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	accessToken, ok := userTokens[token]
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
	}

	var post model.Post

	err = json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = env.FacebookService.Post(post, accessToken)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func handleGetProperties(w http.ResponseWriter, r *http.Request) {
	_, err := extractToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	properties, err := env.PropertyService.GetProperties()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(properties)
}

// Extracts the token from the Authorization header
func extractToken(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", fmt.Errorf("authorization header missing")
	}

	// Expected format: "Bearer <token>"
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", fmt.Errorf("invalid authorization header format")
	}

	return parts[1], nil
}
