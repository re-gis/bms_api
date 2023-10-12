package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/re-gis/bms_api/pkg/db"
	"github.com/re-gis/bms_api/pkg/models"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		loginUser(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed!"))
	}
}

func loginUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	var FoundUser models.User
	// get email and password
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Error while decoding the body", http.StatusInternalServerError)
		return
	}

	if user.Email == "" || user.Password == "" {
		http.Error(w, "All credentials are required!", http.StatusBadRequest)
		return
	}

	if err := db.GetDB().Where("email =?", user.Email).First(&FoundUser).Error; err != nil {
		http.Error(w, "Incorrect credentials provided!", http.StatusBadRequest)
		return
	}

	if FoundUser.Password != user.Password {
		http.Error(w, "Incorrect credentials provided...", http.StatusBadRequest)
		return
	}

	// passed authentication
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(FoundUser)

}
