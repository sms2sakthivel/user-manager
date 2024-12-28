package users

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/sms2sakthivel/user-manager/users/database"
	"github.com/sms2sakthivel/user-manager/users/model"
	"github.com/sms2sakthivel/user-manager/users/repository"
	"github.com/sms2sakthivel/user-manager/users/routes"
	"github.com/sms2sakthivel/user-manager/users/service"
)

func NewApp() *fiber.App {
	// Step 1: Connect to the database
	database.Connect()

	// Step 2: Auto-migrate User schema
	model.Automigrate(database.DB)

	// Step 3: Initialize repository, service, and app
	repo := &repository.GormUserRepository{DB: database.DB}
	service := &service.UserService{Repo: repo}
	app := fiber.New()

	// Step 4: Enable Logger middleware with timing
	app.Use(logger.New(logger.Config{
		Format: "${time} - ${latency} - ${status} - ${method} ${path}\n",
	}))

	// Step 5: Register routes
	routes.RegisterRoutes(app, service)
	return app
}
