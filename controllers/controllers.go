package controllers

import (
	"encoding/json"
	"net/http"

	"go-rest-api/database"
	"go-rest-api/models"

	"github.com/gorilla/mux"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func FindAll(w http.ResponseWriter, r *http.Request) {
	var courses []models.Course
	database.DB.Find(&courses)
	json.NewEncoder(w).Encode(courses)
}

func FindOne(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var course models.Course
	database.DB.First(&course, id)
	json.NewEncoder(w).Encode(course)
}

func Create(w http.ResponseWriter, r *http.Request) {
	var course models.Course
	json.NewDecoder(r.Body).Decode(&course)
	database.DB.Create(&course)
	w.WriteHeader(http.StatusCreated)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var course models.Course
	database.DB.Delete(&course, id)
}

func Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var course models.Course
	database.DB.First(&course, id)
	json.NewDecoder(r.Body).Decode(&course)
	database.DB.Save(&course)
	json.NewEncoder(w).Encode(course)
}
