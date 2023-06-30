package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	// "github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/lukeshay/pms/pkg/adapters"
	"github.com/lukeshay/pms/pkg/auth"
	"github.com/lukeshay/pms/pkg/controllers"
	"github.com/lukeshay/pms/pkg/httputil"
	"github.com/lukeshay/pms/pkg/repositories"

	docs "github.com/lukeshay/pms/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	dsn       = flag.String("dsn", os.Getenv("DATABASE_URL"), "datasource name")
	addr      = flag.String("addr", ":3000", "bind address")
	jwtSecret = flag.String("jwt-secret", os.Getenv("JWT_SECRET"), "jwt secret")
	dist      = flag.String("dist", "frontend-dist", "frontend dist directory")
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

// @title           PMS API
// @version         1.0
// @description     This API contains CRUD operations for PMS.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

//	@securityDefinitions.apikey	ApiKeyAuth
//	@in                         header
//	@name                       Authorization
//	@description                Bearer token for access control

// @BasePath  /api
func main() {
	flag.Parse()

	if *dsn == "" {
		log.Fatalln("dsn required")
		return
	} else if *addr == "" {
		log.Fatalln("addr required")
		return
	} else if *jwtSecret == "" {
		log.Fatalln("jwt-secret required")
		return
	} else if *dist == "" {
		log.Fatalln("dist required")
		return
	}

	environment, present := os.LookupEnv("ENVIRONMENT")
	if !present {
		environment = "production"
	}

	r := gin.Default()

	db := adapters.GetDB(*dsn)
	auther := &auth.Auth{
		JWTSecret:     *jwtSecret,
		SigningMethod: auth.SigningMethodHMAC,
	}
	userRepository := repositories.NewUserRepository(db)
	bookRepository := repositories.NewBookRepository(db)

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

	// ============================= Start of routes =============================

	if environment != "development" {
		gin.SetMode(gin.ReleaseMode)
		r.LoadHTMLGlob(fmt.Sprintf("%s/*.html", *dist))
		r.Use(func(ctx *gin.Context) {
			if strings.HasPrefix(ctx.Request.URL.Path, "/assets") {
				ctx.Header("Cache-Control", "max-age=604800000")
			}

			ctx.Next()
		})
		// r.Use(static.Serve("/", static.LocalFile(*dist, true)))
		r.Static("/assets", fmt.Sprintf("%s/assets", *dist))
		r.NoRoute(func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.html", gin.H{})
		})
	}

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

	docs.SwaggerInfo.BasePath = "/api/v1"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	log.Fatalln(r.Run(*addr))
}
