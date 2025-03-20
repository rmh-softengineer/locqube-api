package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
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
	r.HandleFunc("/auth/login", handleLoginToFacebook).Methods("POST")
	r.HandleFunc("/share-fb-post", handlePostToFacebook).Methods("POST")
	r.HandleFunc("/properties", handleGetProperties).Methods("GET")

	// CORS middleware
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),                             // Allow all origins (change this for security)
		handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"}),        // Allowed HTTP methods
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}), // Allowed headers
	)

	facebookAppID := "640002862073324"

	facebookAppSecret := "6431655c5665aae0c820d90da9666c4c"

	facebookClient := facebookClient.NewClient(facebookAppID, facebookAppSecret)

	facebookService := facebookService.NewService(facebookClient)

	propertyRepository := propertyRepository.NewRepository()

	propertyService := propertyService.NewService(propertyRepository)

	env = appEnv{
		FacebookService: facebookService,
		PropertyService: propertyService,
	}

	fmt.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", corsHandler(r)))
}
