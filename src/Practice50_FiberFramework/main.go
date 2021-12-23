package main

import (
	"Practice50_FiberFramework/book"

	"github.com/gofiber/fiber"
)

// func main() {
//   app := fiber.New()

//   app.Get("/", func(c *fiber.Ctx) {
//     c.Send("Hello, World!")
//   })

//   app.Listen(3000)
// }

func helloWorld(c *fiber.Ctx) {
	c.Send("Hello, World!")
}

func setupRoutes(app *fiber.App) {
	app.Get("/", helloWorld)

	app.Get("/api/v1/book", book.GetBooks)
	app.Get("/api/v1/book/:id", book.GetBook)
	app.Post("/api/v1/book", book.NewBook)
	app.Delete("/api/v1/book/:id", book.DeleteBook)
}

func main() {
	app := fiber.New()

	setupRoutes(app)
	app.Listen(3000)
}
