package server

import (
	"github.com/ThisIsHyum/OpenScheduleApi/internal/repository"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	lg "github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/sirupsen/logrus"
)

func Register(app *fiber.App,
	collegeRepo repository.CollegeRepo, campusRepo repository.CampusRepo,
	studentGroupRepo repository.StudentGroupRepo, callRepo repository.CallRepo,
	lessonRepo repository.LessonRepo,
	createTx repository.CreateTx,
	logger *logrus.Logger, adminToken string) {
	app.Use(cors.New())
	app.Use(lg.New())

	NewCollegeHandler(app, collegeRepo, campusRepo, logger)
	NewCampusHandler(app, campusRepo, studentGroupRepo, collegeRepo, logger)
	NewGroupHandler(app, studentGroupRepo, campusRepo, collegeRepo, logger)
	NewScheduleHandler(app, studentGroupRepo, lessonRepo, callRepo, collegeRepo, logger)
	NewParserHandler(app, callRepo, studentGroupRepo, lessonRepo, campusRepo, collegeRepo, logger)
	NewAdminHandler(app, collegeRepo, campusRepo, createTx, adminToken, logger)
}
