package routers

import (
	"github.com/FernandoCagale/c4-customer/api/handlers"
	"github.com/gorilla/mux"
)

type SystemRoutes struct {
	healthHandler   *handlers.HealthHandler
	customerHandler *handlers.CustomerHandler
}

func (routes *SystemRoutes) MakeHandlers() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/health", routes.healthHandler.Health).Methods("GET")
	r.HandleFunc("/customers", routes.customerHandler.Create).Methods("POST")
	r.HandleFunc("/customers", routes.customerHandler.FindAll).Methods("GET")
	r.HandleFunc("/customers/{id}", routes.customerHandler.FindById).Methods("GET")
	r.HandleFunc("/customers/{id}", routes.customerHandler.DeleteById).Methods("DELETE")

	return r
}

func NewSystem(healthHandler *handlers.HealthHandler, customerHandler *handlers.CustomerHandler) *SystemRoutes {
	return &SystemRoutes{
		healthHandler:   healthHandler,
		customerHandler: customerHandler,
	}
}
