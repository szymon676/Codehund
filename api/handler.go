package api

import (
	"github.com/szymon676/codehund/auth"
	"github.com/szymon676/codehund/service"
)

type Handler struct {
	svc service.Servicer
	sm  *auth.SessionManager
}

func NewHandler(svc service.Servicer, sm *auth.SessionManager) *Handler {
	return &Handler{
		svc: svc,
		sm:  sm,
	}
}
