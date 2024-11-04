package handler

import (
	"20241104/class/2/model"
	"20241104/class/2/repository"
	"20241104/class/2/service"
	"20241104/database"
	"encoding/json"
	"net/http"
)

func UserRegistration(w http.ResponseWriter, r *http.Request) {
	db := database.DbOpen()
	defer db.Close()

	user := model.User{}

	json.NewDecoder(r.Body).Decode(&user)

	result := service.InitUserService(*repository.InitUserRepo(db)).Create(user)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(result)
}
