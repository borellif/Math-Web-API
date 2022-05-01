package main

import (
	"fmt"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/template/html"
)

func main() {
	fmt.Println("Hello world!")

	// Sets location of html files
	engine := html.New("./views", ".html")

	// Attaches the engine to the fiber config views
	app := fiber.New(fiber.Config{
		Views:       engine,
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	// Initializes default CSRF configs and default CORS configs
	app.Use(csrf.New()).Use(cors.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "Math API",
		})
	})

	app.Get("/min", func(c *fiber.Ctx) error {
		c.Request()
	})

	// app.Get("/max", func(c *fiber.Ctx) error {
	// })

	// app.Get("/avg", func(c *fiber.Ctx) error {
	// })

	// app.Get("/median", func(c *fiber.Ctx) error {
	// })

	// app.Get("/percentile", func(c *fiber.Ctx) error {
	// })

	app.Listen(":3000")
}
