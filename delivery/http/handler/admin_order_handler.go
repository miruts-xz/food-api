package handler

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/miruts/food-api/entity"
	"github.com/miruts/food-api/order"
	"net/http"
	"strconv"
)

// AdminUserHandler handles comment related http requests
type AdminOrderHandler struct {
	orderService order.OrderService
}

// NewAdminCommentHandler returns new AdminUserHandler object
func NewAdminOrdertHandler(orderService order.OrderService) *AdminOrderHandler {
	return &AdminOrderHandler{orderService: orderService}
}

// GetComments handles GET /v1/admin/comments request
func (aoh *AdminOrderHandler) GetOrders(w http.ResponseWriter,
	r *http.Request, _ httprouter.Params) {

	var apiKey = r.Header.Get("api-key")
	if apiKey == "" || (apiKey != adminApiKey && apiKey != userApiKey) {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
		return
	}

	orders, errs := aoh.orderService.Orders()

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(orders, "", "\t")

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
func (aoh *AdminOrderHandler) GetSingleOrder(w http.ResponseWriter,
	r *http.Request, ps httprouter.Params) {
	var apiKey = r.Header.Get("api-key")
	if apiKey == "" || (apiKey != userApiKey && apiKey != adminApiKey) {
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

	order, errs := aoh.orderService.Order(uint(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(order, "", "\t")

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
func (aoh *AdminOrderHandler) PostOrder(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	var apiKey = r.Header.Get("api-key")
	if apiKey == "" || (apiKey != userApiKey && apiKey != adminApiKey) {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
		return
	}
	l := r.ContentLength
	body := make([]byte, l)
	r.Body.Read(body)
	order := &entity.Order{}

	err := json.Unmarshal(body, order)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	order, errs := aoh.orderService.StoreOrder(order)

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	output, err := json.MarshalIndent(order, "", "\t")

	w.WriteHeader(http.StatusCreated)
	w.Write(output)
	return
}

// PutComment handles PUT /v1/admin/comments/:id request
func (aoh *AdminOrderHandler) PutOrder(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

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

	order, errs := aoh.orderService.Order(uint(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	l := r.ContentLength

	body := make([]byte, l)

	r.Body.Read(body)

	json.Unmarshal(body, &order)

	order, errs = aoh.orderService.UpdateOrder(order)

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(order, "", "\t")

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
func (aoh *AdminOrderHandler) DeleteOrder(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

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

	_, errs := aoh.orderService.DeleteOrder(uint(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	return
}
