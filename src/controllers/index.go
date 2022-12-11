package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Miguelburitica/goHttpServer/src/models"
)

func ShowHome(w http.ResponseWriter, r *http.Request) {
	colombian := []models.Human{
		{Name: "Miguel", Surname: "Buriticá"},
		{Name: "Andrea", Surname: "Arteaga"},
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(colombian)
}