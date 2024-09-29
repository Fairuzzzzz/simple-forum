package main

import (
	"log"

	"github.com/Fairuzzzzz/simpleform/internal/configs"
	"github.com/Fairuzzzzz/simpleform/internal/handler/membership"
	membershipsRepo "github.com/Fairuzzzzz/simpleform/internal/repository/memberships"
	"github.com/Fairuzzzzz/simpleform/pkg/internalsql"
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

	db, err := internalsql.Connect(cfg.Database.DataSourceName)
	if err != nil {
		log.Fatal("Gagal inisasi database", err)
	}

	_ = membershipsRepo.NewRepository(db)

	membershipHandler := membership.NewHandler(r)
	membershipHandler.RegisterRoute()
	r.Run(cfg.Service.Port)
}
