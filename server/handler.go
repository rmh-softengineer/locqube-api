package server

import (
	"encoding/json"
	"net/http"

	"github.com/rmh-softengineer/locqube/api/model"
)

// In-memory storage for tokens
var userTokens = map[string]string{}

func handleFacebookLogin(w http.ResponseWriter, r *http.Request) {
	authURL := env.FacebookService.GetLoginUrl()
	http.Redirect(w, r, authURL, http.StatusSeeOther)
}

func handleFacebookCallback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "Missing authorization code", http.StatusBadRequest)
		return
	}

	accessToken, err := env.FacebookService.Login(code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userID := r.Header.Get("X-User-ID")
	userTokens[userID] = accessToken

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"access_token": accessToken})
}

func handlePostToFacebook(w http.ResponseWriter, r *http.Request) {
	var post model.Post

	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userID := r.Header.Get("X-User-ID")
	accessToken, ok := userTokens[userID]
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	}

	err = env.FacebookService.Post(post, accessToken)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func handleGetProperties(w http.ResponseWriter, r *http.Request) {
	properties, err := env.PropertyService.GetProperties()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(properties)
}
