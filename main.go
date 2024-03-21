package main

import (
	controller "martini/controller"
	"martini/model"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
)

func main() {
	m := martini.Classic()
	m.Use(render.Renderer())
	m.Use(controller.Connect())

	m.Get("/users", controller.GetAllUsers)
	m.Post("/users", binding.Json(model.User{}), controller.NewUser)
	m.Put("/users", binding.Json(model.User{}), controller.EditUser)
	m.Delete("/users", binding.Json(model.User{}), controller.DeleteUser)

	m.RunOnAddr(":22345")
}
