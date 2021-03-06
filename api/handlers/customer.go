package handlers

import (
	"encoding/json"
	"fmt"
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
		fmt.Println(err.Error())
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
		fmt.Println(err.Error())
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

	if err := handler.usecase.Create(forwardHeaders(r), ecommerce); err != nil {
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

func forwardHeaders(r *http.Request) map[string]string {
	incomingHeaders := []string{
		"x-version",

		// open tracing
		"x-request-id",
		"x-b3-traceid",
		"x-b3-spanid",
		"x-b3-parentspanid",
		"x-b3-sampled",
		"x-b3-flags",
		"x-ot-span-context",
	}

	headers := make(map[string]string)
	for _, th := range incomingHeaders {
		h := r.Header.Get(th)
		if h != "" {
			headers[th] = h
		}
	}

	return headers
}
