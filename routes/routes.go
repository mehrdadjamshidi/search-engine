package routes

import (
	"time"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gofiber/fiber/v2/middleware/cache"
)

// func render(c *fiber.Ctx, component templ.Component, options ...func(*templ.ComponentHandler)) error {
// 	componentHandler := templ.Handler(component)
// 	for _, o := range options {
// 		o(componentHandler)
// 	}
// 	return adaptor.HTTPHandler(componentHandler)(c)
// }

// func Render(c *fiber.Ctx, component templ.Component) error {
// 	c.Set("Content-Type", "text/html")
// 	return component.Render(c.Context(), c.Response().BodyWriter())
// }

// type settingsform struct {
// 	Amount   int  `json:"amount"`
// 	SearchOn bool `json:"search_on"`
// 	AddNew   bool `json:"add_new"`
// }

// func SetRoutes(app *fiber.App) {
// 	// app.Get("/", func(c *fiber.Ctx) error {
// 	// 	// return c.JSON(fiber.Map{
// 	// 	// 	"message": "Hello world",
// 	// 	// })
// 	// 	// return c.SendString("I'm a GET request!")
// 	// 	return render(c, views.Home())
// 	// })
// 	app.Get("/", AuthMiddleware, LoginHandler)

// 	// app.Post("/", func(c *fiber.Ctx) error {
// 	// 	input := settingsform{}
// 	// 	if err := c.BodyParser(&input); err != nil {
// 	// 		return c.SendString("<h2>Error: Something went wrong</h2>")
// 	// 	}
// 	// 	fmt.Println(input)
// 	// 	return c.SendStatus(200)
// 	// })
// 	app.Post("/", AuthMiddleware, LoginPostHandler)

// 	// app.Get("/login",func(c *fiber.Ctx) error {
// 	// 	return render(c, views.Login())
// 	// })
// 	app.Get("/login", LoginHandler)

// 	// app.Post("/login", AuthMiddleware ,func(c *fiber.Ctx) error {
// 	// 	input := loginform{}
// 	// 	if err := c.BodyParser(&input); err != nil {
// 	// 		return c.SendString("<h2>Error: Something went wrong</h2>")
// 	// 	}
// 	// 	return c.SendStatus(200)
// 	// })
// 	app.Post("/login", LoginPostHandler)
// }

func render(c *fiber.Ctx, component templ.Component, options ...func(*templ.ComponentHandler)) error {
	componentHandler := templ.Handler(component)
	for _, o := range options {
		o(componentHandler)
	}
	return adaptor.HTTPHandler(componentHandler)(c)
}

func SetRoutes(app *fiber.App) {
	app.Get("/login", LoginHandler)
	app.Post("/login", LoginPostHandler)
	app.Post("/logout", LogoutHandler)

	app.Post("/search", HandleSearch)
	app.Use("/search", cache.New(cache.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.Query("noCache") == "true"
		},
		Expiration:   30 * time.Minute,
		CacheControl: true,
	}))

	// app.Get("/create", func(c *fiber.Ctx) error {
	// 	u := &db.User{}
	// 	u.CreateAdmin()
	// 	return c.SendString("Admin Created")
	// })

	app.Get("/", AuthMiddleware, DashboardHandler)
	app.Post("/", AuthMiddleware, DashboardPostHandler)
}
