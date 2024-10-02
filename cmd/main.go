package main

import (
	"log"

	"github.com/Fairuzzzzz/simpleform/internal/configs"
	"github.com/Fairuzzzzz/simpleform/internal/handler/membership"
	"github.com/Fairuzzzzz/simpleform/internal/handler/posts"
	membershipsRepo "github.com/Fairuzzzzz/simpleform/internal/repository/memberships"
	postRepo "github.com/Fairuzzzzz/simpleform/internal/repository/posts"
	membershipsSvc "github.com/Fairuzzzzz/simpleform/internal/service/memberships"
	postSvc "github.com/Fairuzzzzz/simpleform/internal/service/posts"
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

	db, err := internalsql.Connect(cfg.Database.DataSourceName)
	if err != nil {
		log.Fatal("Gagal inisasi database", err)
	}

	// Middleware default gin
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	membershipRepo := membershipsRepo.NewRepository(db)
	postRepo := postRepo.NewRepository(db)

	membershipService := membershipsSvc.NewService(cfg, membershipRepo)
	postService := postSvc.NewService(cfg, postRepo)

	membershipHandler := membership.NewHandler(r, membershipService)
	membershipHandler.RegisterRoute()

	postHandler := posts.NewHandler(r, postService)
	postHandler.RegisterRoute()

	r.Run(cfg.Service.Port)
}
