package routes

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/sms2sakthivel/user-manager/users/model"
	"github.com/sms2sakthivel/user-manager/users/service"
)

func RegisterRoutes(app *fiber.App, service *service.UserService) {
	app.Get("/", UserServiceInfo)
	app.Get("/swagger/*", swagger.HandlerDefault)
	app.Get("/users", func(c *fiber.Ctx) error { return GetAllUsers(c, service) })
	app.Get("/users/:id", func(c *fiber.Ctx) error { return GetUserByID(c, service) })
	app.Post("/users", func(c *fiber.Ctx) error { return CreateUser(c, service) })
	app.Put("/users/:id", func(c *fiber.Ctx) error { return UpdateUser(c, service) })
	app.Delete("/users/:id", func(c *fiber.Ctx) error { return DeleteUser(c, service) })
}

// UserServiceInfo returns information about the User Service
//
// @Summary      User Service Info
// @Description  Returns basic information about the User Service
// @Tags         General
// @Accept       json
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Router       / [get]
func UserServiceInfo(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "User Service"})
}

// GetAllUsers retrieves all users
//
// @Summary      Get All Users
// @Description  Retrieve a list of all users
// @Tags         Users
// @Accept       json
// @Produce      json
// @Success      200  {array}   model.User
// @Failure      500  {object}  fiber.Error
// @Router       /users [get]
func GetAllUsers(c *fiber.Ctx, service *service.UserService) error {
	users, err := service.GetUsers()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(users)
}

// GetUserByID retrieves a user by their ID
//
// @Summary      Get User by ID
// @Description  Retrieve a user by their ID
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  model.User
// @Failure      404  {object}  fiber.Error
// @Failure      500  {object}  fiber.Error
// @Router       /users/{id} [get]
func GetUserByID(c *fiber.Ctx, service *service.UserService) error {
	id, _ := strconv.Atoi(c.Params("id"))
	user, err := service.GetUser(uint(id))
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "User not found")
	}
	return c.JSON(user)
}

// CreateUser adds a new user
//
// @Summary      Create a New User
// @Description  Add a new user to the system
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        user  body      model.User  true  "User details"
// @Success      201   {object}  model.User
// @Failure      400   {object}  fiber.Error
// @Failure      500   {object}  fiber.Error
// @Router       /users [post]
func CreateUser(c *fiber.Ctx, service *service.UserService) error {
	var user model.User
	if err := c.BodyParser(&user); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid input")
	}
	if err := service.CreateUser(&user); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.Status(fiber.StatusCreated).JSON(user)
}

// UpdateUser modifies details of an existing user
//
// @Summary      Update an Existing User
// @Description  Modify details of an existing user
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        id    path      int         true  "User ID"
// @Param        user  body      model.User  true  "Updated user details"
// @Success      200   {object}  model.User
// @Failure      400   {object}  fiber.Error
// @Failure      404   {object}  fiber.Error
// @Failure      500   {object}  fiber.Error
// @Router       /users/{id} [put]
func UpdateUser(c *fiber.Ctx, service *service.UserService) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var user model.User
	if err := c.BodyParser(&user); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid input")
	}
	user.ID = uint(id)
	if err := service.UpdateUser(&user); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(user)
}

// DeleteUser removes a user by their ID
//
// @Summary      Delete a User
// @Description  Remove a user by their ID
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        id   path  int  true  "User ID"
// @Success      204
// @Failure      500  {object}  fiber.Error
// @Router       /users/{id} [delete]
func DeleteUser(c *fiber.Ctx, service *service.UserService) error {
	id, _ := strconv.Atoi(c.Params("id"))
	if err := service.DeleteUser(uint(id)); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.SendStatus(fiber.StatusNoContent)
}
