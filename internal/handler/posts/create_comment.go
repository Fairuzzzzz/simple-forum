package posts

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Fairuzzzzz/simpleform/internal/model/posts"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateComment(c *gin.Context) {
	ctx := c.Request.Context()

	var request posts.CreateCommentRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Mengambil postID dari path http dan ubah ke bentuk integer
	postIDStr := c.Param("postID")
	postID, err := strconv.ParseInt(postIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errors.New("postID pada param tidak valid").Error(),
		})
		return
	}

	// Mengambil userID dari middleware
	userID := c.GetInt64("userID")

	err = h.postSvc.CreateComment(ctx, postID, userID, request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.Status(http.StatusCreated)
}
