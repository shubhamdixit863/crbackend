package app

import (
	"encoding/json"
	"microservicesgo/service"
	"net/http"
)

type UserHandlers struct {
	service service.UserService
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}

func (ch *UserHandlers) getAllUsers(w http.ResponseWriter, r *http.Request) {

	customers, err := ch.service.GetAllUsers()

	if err != nil {
		writeResponse(w, 400, nil)
	} else {
		writeResponse(w, http.StatusOK, customers)
	}
}
