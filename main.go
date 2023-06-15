package main

import (
	"log"
	"os"

	"github.com/lukeshay/pms/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"

	_ "github.com/lukeshay/pms/docs"
)

// @title           Some API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @BasePath  /api
func main() {
	environment, present := os.LookupEnv("ENVIRONMENT")
	if !present {
		environment = "production"
	}

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Get("/swagger/*", swagger.HandlerDefault) // default

	app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL:         "http://localhost:3000/doc.json",
		DeepLinking: false,
		// Expand ("list") or Collapse ("none") tag groups by default
		DocExpansion: "none",
		// // Prefill OAuth ClientId on Authorize popup
		// OAuth: &swagger.OAuthConfig{
		// 	AppName:  "OAuth Provider",
		// 	ClientId: "21bb4edc-05a7-4afc-86f1-2e151e4ba6e2",
		// },
		// // Ability to change OAuth2 redirect uri location
		// OAuth2RedirectUrl: "http://localhost:8080/swagger/oauth2-redirect.html",
	}))

	v1 := app.Group("/api/v1")

	booksV1Controller := controllers.NewBooksV1Controller()
	booksV1 := v1.Group(booksV1Controller.BasePath())

	booksV1.Get("/", booksV1Controller.List)
	booksV1.Post("/", booksV1Controller.Post)
	booksV1.Get("/{id:string}", booksV1Controller.Get)
	booksV1.Put("/{id:string}", booksV1Controller.Put)
	booksV1.Delete("/{id:string}", booksV1Controller.Delete)

	if environment != "development" {
		app.Static("/", "./frontend-dist")

		app.Get("/*", func(ctx *fiber.Ctx) error {
			return ctx.SendFile("./frontend-dist/index.html")
		})
	}

	log.Fatal(app.Listen(":3000"))
}
