package main

import (
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lukeshay/pms/pkg/adapters"
	"github.com/lukeshay/pms/pkg/auth"
	"github.com/lukeshay/pms/pkg/controllers"
	"github.com/lukeshay/pms/pkg/httputil"
	"github.com/lukeshay/pms/pkg/repositories"

	"github.com/gin-contrib/cors"

	_ "github.com/lukeshay/pms/docs"
)

func RequireAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		claims, err := auth.GetClaims(ctx)

		if err != nil || claims == nil {
			ctx.JSON(http.StatusUnauthorized, httputil.HTTPError{
				Error: httputil.HTTPErrorError{
					Message: "Unauthorized",
				},
			})

			return
		}

		ctx.Next()
	}
}

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

	r := gin.Default()

	db := adapters.GetDB()
	auther := &auth.Auth{
		JWTSecret:     os.Getenv("JWT_SECRET"),
		SigningMethod: auth.SigningMethodHMAC,
	}
	userRepository := repositories.NewUserRepository(db)
	userRepository.CreateTable()
	bookRepository := repositories.NewBookRepository(db)
	bookRepository.CreateTable()

	r.Use(gin.Logger(), gin.Recovery(), cors.New(cors.Config{
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type", "Accept-Encoding"},
		AllowAllOrigins:  true,
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.Use(func(ctx *gin.Context) {
		authorization := strings.Replace(ctx.GetHeader("Authorization"), "Bearer ", "", 1)
		if authorization == "" {
			println("no authorization")
			ctx.Next()
			return
		}

		claims, err := auther.JWTParse(authorization)
		if err != nil {
			println("not parsable", err.Error())
			ctx.Next()
			return
		}

		auth.SetClaims(ctx, claims)
		ctx.Next()
	})

	v1 := r.Group("/api/v1")

	booksV1Controller := controllers.NewBooksV1Controller(controllers.BooksV1ControllerInput{
		BookRepository: bookRepository,
	})
	{
		booksV1 := v1.Group(booksV1Controller.BasePath())
		{
			booksV1.Use(RequireAuth())

			booksV1.GET("/", booksV1Controller.List)
			booksV1.POST("/", booksV1Controller.Post)
			booksV1.GET("/:id/", booksV1Controller.Get)
			booksV1.PUT("/:id/", booksV1Controller.Put)
			booksV1.DELETE("/:id/", booksV1Controller.Delete)
		}
	}

	authV1Controller := controllers.NewAuthV1Controller(controllers.NewAuthV1ControllerInput{
		UserRepository: userRepository,
		Auth:           auther,
	})
	{
		authV1 := v1.Group(authV1Controller.BasePath())
		{
			authV1.GET("/", authV1Controller.Get)
			authV1.POST("/", authV1Controller.Post)
			authV1.POST("/sign-in/", authV1Controller.PostSignIn)
		}
	}

	if environment != "development" {
		gin.SetMode(gin.ReleaseMode)
		r.Static("/assets", "./frontend-dist/assets")
		r.NoRoute(func(c *gin.Context) {
			c.File("./frontend-dist/index.html")
			c.Status(200)
		})
	}

	log.Fatal(r.Run(":3000"))
}
