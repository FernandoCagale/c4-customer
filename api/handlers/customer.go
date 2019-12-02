package handlers

import (
	"encoding/json"
	"github.com/FernandoCagale/c4-customer/api/render"
	"github.com/FernandoCagale/c4-customer/internal/errors"
	"github.com/FernandoCagale/c4-customer/pkg/domain/customer"
	"github.com/FernandoCagale/c4-customer/pkg/entity"
	"github.com/gorilla/mux"
	"net/http"
)

type CustomerHandler struct {
	usecase customer.UseCase
}

func NewCustomer(usecase customer.UseCase) *CustomerHandler {
	return &CustomerHandler{
		usecase: usecase,
	}
}

func (handler *CustomerHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	customers, err := handler.usecase.FindAll()
	if err != nil {
		render.ResponseError(w, err, http.StatusInternalServerError)
		return
	}

	render.Response(w, customers, http.StatusOK)
}

func (handler *CustomerHandler) FindById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	ID := vars["id"]

	customer, err := handler.usecase.FindById(ID)
	if err != nil {
		switch err {
		case errors.ErrNotFound:
			render.ResponseError(w, err, http.StatusNotFound)
		default:
			render.ResponseError(w, err, http.StatusInternalServerError)
		}
		return
	}

	render.Response(w, customer, http.StatusOK)
}

func (handler *CustomerHandler) DeleteById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	ID := vars["id"]

	err := handler.usecase.DeleteById(ID)
	if err != nil {
		switch err {
		case errors.ErrNotFound:
			render.ResponseError(w, err, http.StatusNotFound)
		default:
			render.ResponseError(w, err, http.StatusInternalServerError)
		}
		return
	}

	render.Response(w, nil, http.StatusNoContent)
}

func (handler *CustomerHandler) Create(w http.ResponseWriter, r *http.Request) {
	var ecommerce *entity.Ecommerce

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&ecommerce); err != nil {
		render.ResponseError(w, err, http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	if err := handler.usecase.Create(ecommerce); err != nil {
		switch err {
		case errors.ErrInvalidPayload:
			render.ResponseError(w, err, http.StatusBadRequest)
		case errors.ErrConflict:
			render.ResponseError(w, err, http.StatusConflict)
		default:
			render.ResponseError(w, err, http.StatusInternalServerError)
		}
		return
	}

	render.Response(w, nil, http.StatusCreated)
}
