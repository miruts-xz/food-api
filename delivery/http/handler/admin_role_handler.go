package handler

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/miruts/food-api/entity"
	"github.com/miruts/food-api/user"
	"net/http"
	"strconv"
)

// AdminRoleHandler is used to implement role related http requests
type AdminRoleHandler struct {
	roleService user.RoleService
}

// NewAdminRoleHandler initializes and returns new AdminRoleHandler object
func NewAdminRoleHandler(roleSrv user.RoleService) *AdminRoleHandler {
	return &AdminRoleHandler{roleService: roleSrv}
}

// GetRoles handles GET /v1/admin/roles requests
func (arh *AdminRoleHandler) GetRoles(w http.ResponseWriter,
	r *http.Request, _ httprouter.Params) {

	roles, errs := arh.roleService.Roles()

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(roles, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return

}

// GetSingleRole handles GET /v1/admin/roles/:id requests
func (arh *AdminRoleHandler) GetSingleRole(w http.ResponseWriter,
	r *http.Request, ps httprouter.Params) {

		id, err := strconv.Atoi(ps.ByName("id"))

		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		role, errs := arh.roleService.Role(uint(id))

		if len(errs) > 0 {
			w.Header().Set("Content-Type", "application/json")
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		output, err := json.MarshalIndent(role, "", "\t")

		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(output)
		return
}

// PutRole handles PUT /v1/admin/roles/:id requests
func (arh *AdminRoleHandler) PutRole(w http.ResponseWriter,
	r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	role, errs := arh.roleService.Role(uint(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	l := r.ContentLength

	body := make([]byte, l)

	r.Body.Read(body)

	json.Unmarshal(body, &role)

	role, errs = arh.roleService.UpdateRole(role)

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(role, "", "\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return

}

// PostRole handles POST /v1/admin/roles requests
func (arh *AdminRoleHandler) PostRole(w http.ResponseWriter,
	r *http.Request, _ httprouter.Params) {
	l := r.ContentLength
	body := make([]byte, l)
	r.Body.Read(body)
	role := &entity.Role{}

	err := json.Unmarshal(body, role)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	role, errs := arh.roleService.StoreRole(role)

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	p := fmt.Sprintf("/v1/admin/roles/%d", role.ID)
	w.Header().Set("Location", p)
	w.WriteHeader(http.StatusCreated)
	return

}

// DeleteRole handles DELETE /v1/admin/roles/:id requests
func (arh *AdminRoleHandler) DeleteRole(w http.ResponseWriter,
	r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	_, errs := arh.roleService.DeleteRole(uint(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	return
}
