package routers

import (
	"github.com/FernandoCagale/c4-customer/api/event"
	"github.com/FernandoCagale/c4-customer/api/handlers"
	"github.com/gorilla/mux"
	"time"
)

type SystemRoutes struct {
	healthHandler   *handlers.HealthHandler
	customerHandler *handlers.CustomerHandler
	customerEvent   *event.CustomerEvent
}

func (routes *SystemRoutes) MakeEvents() {
	time.Sleep(5 * time.Second)

	routes.customerEvent.ProcessCustomer()
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

func NewSystem(healthHandler *handlers.HealthHandler, customerHandler *handlers.CustomerHandler, customerEvent *event.CustomerEvent) *SystemRoutes {
	return &SystemRoutes{
		healthHandler:   healthHandler,
		customerHandler: customerHandler,
		customerEvent:   customerEvent,
	}
}
