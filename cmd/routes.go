package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/michpalm/manage-me/user/login"
	"github.com/michpalm/manage-me/user/register"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", register.Home)

	app.Get("/register", register.NewView)
	app.Post("/register", register.Register)
	app.Get("/login", login.NewView)
	app.Post("/login", login.Login)
}
