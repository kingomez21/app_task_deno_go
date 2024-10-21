package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kingomez21/app_task_deno_go/controllers"
)

func UserRoute(app *fiber.App) fiber.Router {

	user := app.Group("/users")

	user.Get("", controllers.GetUsers)
	user.Post("", controllers.CreateUser)
	user.Get("/:id", controllers.GetOneUser)
	user.Put("/:id", controllers.UpdatedUser)
	user.Delete("/:id", controllers.DeleteUser)

	return user
}
