package membership

import (
	"net/http"

	"github.com/Fairuzzzzz/simpleform/internal/model/memberships"
	"github.com/gin-gonic/gin"
)

func (h *Handler) Refresh(c *gin.Context) {
	ctx := c.Request.Context()

	var request memberships.RefreshTokenRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userID := c.GetInt64("userID")

	accessToken, err := h.membershipSvc.ValidateRefreshToken(ctx, userID, request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, memberships.RefreshResponse{
		AccessToken: accessToken,
	})
}
