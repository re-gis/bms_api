package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/re-gis/bms_api/pkg/db"
	"github.com/re-gis/bms_api/pkg/models"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		registerUser(w, r)
	case http.MethodGet:
		getUsers(w, r)
	case http.MethodPut:
		// put
	case http.MethodDelete:
		// delete
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed!"))
	}
}

func registerUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Error while decoding the body!", http.StatusInternalServerError)
		return
	}

	if user.Email == "" || user.Password == "" || user.Name == "" {
		http.Error(w, "All credentials are required!", http.StatusBadRequest)
		return
	}

	result := db.GetDB().Create(&user)
	if result.Error != nil {
		http.Error(w, "Error while creating the user...", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User

	result := db.GetDB().Find(&users)
	if result.Error != nil {
		http.Error(w, "An error occurred while fetching users...", http.StatusInternalServerError)
		return
	}

	jsonRes, err := json.Marshal(users)
	if err != nil {
		http.Error(w, "An error occurred while parsing users...", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonRes)
}

func getOneUser(w http.ResponseWriter, r *http.Request) {

}
