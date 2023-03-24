package transport

import "league/internal/services"

type Handler struct {
	Services *services.Services
}

func New(services *services.Services) *Handler {
	return &Handler{
		Services: services,
	}
}
