package handler

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/miruts/food-api/entity"
	"github.com/miruts/food-api/user"
	"net/http"
	"strconv"
)

// AdminUserHandler handles comment related http requests
type AdminUserHandler struct {
	userService user.UserService
}

// NewAdminCommentHandler returns new AdminUserHandler object
func NewAdminUserHandler(userService user.UserService) *AdminUserHandler {
	return &AdminUserHandler{userService: userService}
}

// GetComments handles GET /v1/admin/comments request
func (auh *AdminUserHandler) GetUsers(w http.ResponseWriter,
	r *http.Request, _ httprouter.Params) {

	var apiKey = r.Header.Get("api-key")
	if apiKey == "" || apiKey != adminApiKey {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
		return
	}
	users, errs := auh.userService.Users()

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(users, "", "\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return

}

// GetSingleComment handles GET /v1/admin/comments/:id request
func (auh *AdminUserHandler) GetSingleUser(w http.ResponseWriter,
	r *http.Request, ps httprouter.Params) {

	var apiKey = r.Header.Get("api-key")
	if apiKey == "" || (apiKey != adminApiKey && apiKey != userApiKey) {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
		return
	}

	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	user, errs := auh.userService.User(uint(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(user, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

// PostComment handles POST /v1/admin/comments request
func (auh *AdminUserHandler) PostUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var apiKey = r.Header.Get("api-key")
	if apiKey == "" || (apiKey != adminApiKey && apiKey != userApiKey) {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
		return
	}
	l := r.ContentLength
	body := make([]byte, l)
	r.Body.Read(body)
	user := &entity.User{}

	err := json.Unmarshal(body, user)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	user, errs := auh.userService.StoreUser(user)

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusCreated)
	return
}

// PutComment handles PUT /v1/admin/comments/:id request
func (auh *AdminUserHandler) PutUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	var apiKey = r.Header.Get("api-key")
	if apiKey == "" || (apiKey != adminApiKey && apiKey != userApiKey) {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
		return
	}
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	user, errs := auh.userService.User(uint(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	l := r.ContentLength

	body := make([]byte, l)

	r.Body.Read(body)

	json.Unmarshal(body, &user)

	user, errs = auh.userService.UpdateUser(user)

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(user, "", "\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

// DeleteComment handles DELETE /v1/admin/comments/:id request
func (auh *AdminUserHandler) DeleteUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	var apiKey = r.Header.Get("api-key")
	if apiKey == "" || (apiKey != adminApiKey && apiKey != userApiKey) {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
		return
	}
	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	_, errs := auh.userService.DeleteUser(uint(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	return
}
func (auh *AdminUserHandler) GetByUsername(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var apiKey = r.Header.Get("api-key")
	if apiKey == "" || (apiKey != adminApiKey && apiKey != userApiKey) {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
		return
	}
	username := ps.ByName("username")

	usr, errs := auh.userService.UserByUsername(username)

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(usr, "", "\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}
