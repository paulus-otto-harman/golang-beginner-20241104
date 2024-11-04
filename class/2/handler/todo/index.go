package todo

import (
	"20241104/class/2/model"
	"20241104/class/2/repository"
	"20241104/class/2/service"
	"20241104/database"
	"encoding/json"
	"net/http"
)

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	db := database.DbOpen()
	defer db.Close()

	result := service.InitTodoService(*repository.InitTodoRepo(db)).Get(model.Session{Id: r.Header.Get("token")})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}
