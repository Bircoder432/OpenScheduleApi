package main

import (
	"fmt"

	"github.com/ThisIsHyum/OpenScheduleApi/internal/config"
	"github.com/ThisIsHyum/OpenScheduleApi/internal/database"
	"github.com/ThisIsHyum/OpenScheduleApi/internal/logger"
	"github.com/ThisIsHyum/OpenScheduleApi/internal/server"
	"github.com/gofiber/fiber/v3"
)

func main() {
	logger := logger.New()

	cfg, err := config.LoadConfig()
	if err != nil {
		logger.WithError(err).Fatal("unable to load config")
	}

	db, err := database.NewDb(&cfg.Db)
	if err != nil {
		logger.WithError(err).Fatal("unable to connect database")
	}

	collegeRepo := database.NewCollegeDb(db)
	campusRepo := database.NewCampusDb(db)
	studentGroupRepo := database.NewGroupDb(db)
	callRepo := database.NewCallDb(db)
	lessonRepo := database.NewLessonDb(db)

	createTx := database.InitCreateTx(db)

	app := fiber.New()
	server.Register(app,
		collegeRepo, campusRepo,
		studentGroupRepo, callRepo, lessonRepo, createTx, logger, cfg.AdminToken)

	addr := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
	logger.Infof("Running server on %s", addr)
	if err := app.Listen(addr, fiber.ListenConfig{
		DisableStartupMessage: true,
	}); err != nil {
		logger.WithError(err).Fatal("unable to run server")
	}
}
