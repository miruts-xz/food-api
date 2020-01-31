package handler

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/miruts/food-api/entity"
	"github.com/miruts/food-api/menu"
	"net/http"
	"strconv"
)

// AdminUserHandler handles comment related http requests
type AdminItemHandler struct {
	itemService menu.ItemService
}

// NewAdminCommentHandler returns new AdminUserHandler object
func NewAdminItemHandler(itemService menu.ItemService) *AdminItemHandler {
	return &AdminItemHandler{itemService: itemService}
}

// GetComments handles GET /v1/admin/comments request
func (aingh *AdminItemHandler) GetItems(w http.ResponseWriter,
	r *http.Request, _ httprouter.Params) {

	items, errs := aingh.itemService.Items()

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(items, "", "\t")

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
func (aingh *AdminItemHandler) GetSingleItem(w http.ResponseWriter,
	r *http.Request, ps httprouter.Params) {

	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	item, errs := aingh.itemService.Item(uint(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(item, "", "\t")

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
func (aingh *AdminItemHandler) PostItem(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	l := r.ContentLength
	body := make([]byte, l)
	r.Body.Read(body)
	item := &entity.Item{}

	err := json.Unmarshal(body, item)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	item, errs := aingh.itemService.StoreItem(item)

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	p := fmt.Sprintf("/v1/admin/items/%d", item.ID)
	w.Header().Set("Location", p)
	w.WriteHeader(http.StatusCreated)
	return
}

// PutComment handles PUT /v1/admin/comments/:id request
func (aingh *AdminItemHandler) PutItems(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	item, errs := aingh.itemService.Item(uint(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	l := r.ContentLength

	body := make([]byte, l)

	r.Body.Read(body)

	json.Unmarshal(body, &item)

	item, errs = aingh.itemService.UpdateItem(item)

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(item, "", "\t")

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
func (aingh *AdminItemHandler) DeleteItem(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	_, errs := aingh.itemService.DeleteItem(uint(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	return
}
