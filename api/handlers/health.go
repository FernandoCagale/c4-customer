package handlers

import (
	"github.com/hellofresh/health-go"
	"net/http"
)

type HealthHandler struct {
}

func NewHealth() *HealthHandler {
	return &HealthHandler{}
}

func (handler *HealthHandler) Health(w http.ResponseWriter, r *http.Request) {
	//health.Register(health.Config{
	//	Name:    "rabbitmq",
	//	Timeout: time.Second * 5,
	//	Check: healthRabbitmq.New(healthRabbitmq.Config{
	//		DSN: os.Getenv("AMQP_URI"),
	//	}),
	//})
	//
	//health.Register(health.Config{
	//	Name:    "postgresql",
	//	Timeout: time.Second * 5,
	//	Check: healthPostgres.New(healthPostgres.Config{
	//		DSN: os.Getenv("POSTGRES_ADDRS"),
	//	}),
	//})

	health.HandlerFunc(w, r)
}
