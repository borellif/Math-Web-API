package main

import (
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/template/html"

	"github.com/borellif/Math-Web-Api/math"
)

func main() {
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

	setupRoutes(app)
	app.Listen(":3000")
}

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/min", math.Minimum)
	app.Get("/api/v1/max", math.Maximum)
	app.Get("/api/v1/avg", math.Average)
	app.Get("/api/v1/median", math.Median)
	app.Get("/api/v1/percentile", math.Percentile)
}
