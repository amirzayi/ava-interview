package api

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/amirzayi/ava-interview/pkg/jsonutil"
)

var ErrInvalidIDParameter = errors.New("invalid id parameter")

func (a *Router) ListUsers(w http.ResponseWriter, r *http.Request) {
	users, err := a.Service.ListUsers(r.Context())
	if err != nil {
		jsonutil.Encode(w, http.StatusInternalServerError, err.Error())
		return
	}

	jsonutil.Encode(w, http.StatusOK, ModelsToDTOs(users))
}

func (a *Router) CreateUser(w http.ResponseWriter, r *http.Request) {
	req, err := jsonutil.Decode[User](r)
	if err != nil {
		jsonutil.Encode(w, http.StatusBadRequest, err.Error())
		return
	}

	user, err := a.Service.CreateUser(r.Context(), req.CreateParam())
	if err != nil {
		jsonutil.Encode(w, http.StatusInternalServerError, err.Error())
		return
	}
	jsonutil.Encode(w, http.StatusCreated, ModelToDTO(user))
}

func (a *Router) GetUserByID(w http.ResponseWriter, r *http.Request) {
	idParam := r.PathValue("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		http.Error(w, ErrInvalidIDParameter.Error(), http.StatusBadRequest)
		return
	}

	user, err := a.Service.GetUserByID(r.Context(), id)
	if err != nil {
		jsonutil.Encode(w, http.StatusInternalServerError, err.Error())
		return
	}

	jsonutil.Encode(w, http.StatusOK, ModelToDTO(user))
}

func (a *Router) DeleteUserByID(w http.ResponseWriter, r *http.Request) {
	idParam := r.PathValue("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		http.Error(w, ErrInvalidIDParameter.Error(), http.StatusBadRequest)
		return
	}

	err = a.Service.DeleteUserByID(r.Context(), id)
	if err != nil {
		jsonutil.Encode(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
