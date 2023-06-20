package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/lukeshay/pms/pkg/adapters"
	"github.com/lukeshay/pms/pkg/auth"
	"github.com/lukeshay/pms/pkg/controllers"
	"github.com/lukeshay/pms/pkg/httputil"
	"github.com/lukeshay/pms/pkg/repositories"

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

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// @title           Some API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

//	@securityDefinitions.apikey	ApiKeyAuth
//	@in                         header
//	@name                       Authorization
//	@description                Description for what is this security definition being used

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

	r.Use(gin.Logger(), gin.Recovery(), Cors())

	r.Use(func(ctx *gin.Context) {
		authorization := ctx.GetHeader("Authorization")
		if authorization == "" {
			println("no authorization, checking cookies")
			var err error
			authorization, err = ctx.Cookie("Authorization")

			if err != nil {
				println("no cookie", err.Error())
				ctx.Next()
				return
			}
		}

		authorization = strings.Replace(authorization, "Bearer ", "", 1)

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

// /api/v1/auth/sign-in/
// /api/v1/auth/sign-in/
