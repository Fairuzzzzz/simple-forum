package main

import (
	"log"

	"github.com/Fairuzzzzz/simpleform/internal/configs"
	"github.com/Fairuzzzzz/simpleform/internal/handler/membership"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	var (
		cfg *configs.Config
	)

	err := configs.Init(
		configs.WithConfigFolder([]string{"./internal/configs/"}),
		configs.WithConfigFile("config"),
		configs.WithConfigType("yaml"),
	)

	if err != nil {
		log.Fatal("Gagal inisialisasi config", err)
	}

	cfg = configs.Get()
	log.Println("config", cfg)

	membershipHandler := membership.NewHandler(r)
	membershipHandler.RegisterRoute()
	r.Run(cfg.Service.Port)
}
