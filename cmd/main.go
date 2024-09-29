package main

import (
	"github.com/Fairuzzzzz/simpleform/internal/handler/membership"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	membershipHandler := membership.NewHandler(r)
	membershipHandler.RegisterRoute()
	r.Run(":8000")
}
