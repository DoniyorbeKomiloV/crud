package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"user/models"
)

func (h *Handler) User(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		h.CreateUser(w, r)
	}
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var createUser models.CreateUser

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error while creating user"))
		return
	}
	err = json.Unmarshal(body, &createUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error while unmarshalling body"))
		return
	}
	userId, err := h.strg.User().Create(&createUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error while create user"))
		return
	}
	user, err := h.strg.User().GetById(&models.UserPrimaryKey{
		Id: userId.Id,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error server getbyid user"))
		return
	}
	resp, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error while marshaling resp"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(resp)

}
