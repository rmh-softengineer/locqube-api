package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	facebookClient "github.com/rmh-softengineer/locqube/api/http/facebook"
	propertyRepository "github.com/rmh-softengineer/locqube/api/repository/property"
	facebookService "github.com/rmh-softengineer/locqube/api/service/facebook"
	propertyService "github.com/rmh-softengineer/locqube/api/service/property"
)

var env appEnv

type appEnv struct {
	FacebookService *facebookService.Service
	PropertyService *propertyService.PropertyService
}

func StartServer() {
	r := mux.NewRouter()
	r.HandleFunc("/auth/login", handleFacebookLogin)
	r.HandleFunc("/auth/callback", handleFacebookCallback)
	r.HandleFunc("/post", postToFacebook).Methods("POST")

	facebookAppID := ""

	facebookAppSecret := ""

	redirectURL := "http://localhost:8080/auth/callback"

	facebookClient := facebookClient.NewClient(facebookAppID, facebookAppSecret, redirectURL)

	facebookService := facebookService.NewService(facebookClient)

	propertyRepository := propertyRepository.NewRepository()

	propertyService := propertyService.NewService(propertyRepository)

	env = appEnv{
		FacebookService: facebookService,
		PropertyService: propertyService,
	}

	fmt.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
