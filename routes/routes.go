package routes

import (
	"log"
	"net/http"

	"go-rest-api/controllers"
	"go-rest-api/middlewares"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middlewares.SetHeaders)
	r.HandleFunc("/api/v1/health", controllers.HealthCheck)
	r.HandleFunc("/api/v1/courses", controllers.FindAll).Methods("GET")
	r.HandleFunc("/api/v1/courses/{id}", controllers.FindOne).Methods("GET")
	r.HandleFunc("/api/v1/courses", controllers.Create).Methods("POST")
	r.HandleFunc("/api/v1/courses/{id}", controllers.Delete).Methods("DELETE")
	r.HandleFunc("/api/v1/courses/{id}", controllers.Update).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
