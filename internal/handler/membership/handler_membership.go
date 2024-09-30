package membership

import (
	"context"

	"github.com/Fairuzzzzz/simpleform/internal/model/memberships"
	"github.com/gin-gonic/gin"
)

type membershipService interface {
	SignUp(ctx context.Context, req memberships.SignUpRequest) error
	Login(ctx context.Context, req memberships.LoginRequest) (string, error)
}

type Handler struct {
	*gin.Engine

	membershipSvc membershipService
}

func NewHandler(api *gin.Engine, membershipSvc membershipService) *Handler {
	return &Handler{
		Engine:        api,
		membershipSvc: membershipSvc,
	}
}

func (h *Handler) RegisterRoute() {
	route := h.Group("membership")
	route.GET("/ping", h.Ping)
	route.POST("/sign-up", h.SignUp)
	route.POST("/login", h.Login)
}
