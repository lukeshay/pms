package controllers

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lukeshay/pms/pkg/auth"
	"github.com/lukeshay/pms/pkg/httputil"
	"github.com/lukeshay/pms/pkg/models"
	"github.com/lukeshay/pms/pkg/repositories"
)

type AuthV1Controller struct {
	basePath       string
	userRepository *repositories.UserRepository
	auth           *auth.Auth
}

type NewAuthV1ControllerInput struct {
	UserRepository *repositories.UserRepository
	Auth           *auth.Auth
}

func NewAuthV1Controller(input NewAuthV1ControllerInput) *AuthV1Controller {
	return &AuthV1Controller{
		basePath:       "/auth",
		userRepository: input.UserRepository,
		auth:           input.Auth,
	}
}

func (c *AuthV1Controller) BasePath() string {
	return c.basePath
}

type AuthGetResponse struct {
	Claims auth.Claims `json:"claims" binding:"required"`
}

// AuthGet godoc
// @Summary      Get jwt claims
// @Description  get jwt claims
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Success      200  {object}  AuthGetResponse
// @Failure      404  {object}  httputil.HTTPError
// @Router       /v1/auth/ [get]
func (c *AuthV1Controller) Get(ctx *gin.Context) {
	claims, err := auth.GetClaims(ctx)
	if !httputil.RespondInternalServerErrorQuery(ctx, "Unknown error", claims, err) {
		return
	}

	httputil.RespondOK(ctx, AuthGetResponse{
		Claims: *claims,
	})
}

type AuthPostRequest struct {
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
}

type AuthTokens struct {
	Authorization string `json:"authorization" binding:"required"`
}

type AuthPostResponse struct {
	Tokens AuthTokens  `json:"tokens" binding:"required"`
	User   models.User `json:"user" binding:"required"`
}

// AuthPost godoc
// @Summary      Create user
// @Description  create user
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param			   user	body		  AuthPostRequest	    true  "Create user"
// @Success      200  {object}  AuthPostResponse
// @Failure      400  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /v1/auth/ [post]
func (c *AuthV1Controller) Post(ctx *gin.Context) {
	var request AuthPostRequest
	if !httputil.ShouldBindJSON(ctx, &request) {
		return
	}

	hashedPassword, err := c.auth.PasswordHash(request.Password)
	if !httputil.RespondInternalServerErrorQuery(ctx, "Unknown error", &hashedPassword, err) {
		return
	}

	userId := uuid.NewString()

	user := &models.User{
		Email:         request.Email,
		EmailVerified: false,
		FirstName:     request.FirstName,
		Id:            userId,
		LastName:      request.LastName,
		Password:      hashedPassword,
		Timestamps: models.Timestamps{
			CreatedAt: time.Now(),
			CreatedBy: userId,
			UpdatedAt: time.Now(),
			UpdatedBy: userId,
		},
	}

	err = c.userRepository.Insert(user)
	if !httputil.RespondInternalServerErrorQuery(ctx, "Unknown error", user, err) {
		return
	}

	jwt, err := c.auth.JWTGenerate(user)
	if !httputil.RespondInternalServerErrorQuery(ctx, "Unknown error", &jwt, err) {
		return
	}

	httputil.RespondOK(ctx, AuthPostResponse{
		Tokens: AuthTokens{
			Authorization: jwt,
		},
		User: *user,
	})
}

type AuthPostSignInRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// AuthPostSignIn godoc
// @Summary      Sign in to user
// @Description  sign in to user
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param			   user	body		  AuthPostSignInRequest  true  "Sign in to user"
// @Success      200  {object}  AuthPostResponse
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /v1/auth/sign-in [post]
func (c *AuthV1Controller) PostSignIn(ctx *gin.Context) {
	var request AuthPostSignInRequest

	if !httputil.ShouldBindJSON(ctx, &request) {
		return
	}

	user, err := c.userRepository.GetByEmail(request.Email)
	if !httputil.RespondNotFoundQuery(ctx, "User not found", user, err) {
		return
	}

	if !c.auth.PasswordHashCheck(request.Password, user.Password) {
		httputil.RespondNotFound(ctx, "User not found")
		return
	}

	jwt, err := c.auth.JWTGenerate(user)
	if !httputil.RespondInternalServerErrorQuery(ctx, "Unknown error", &jwt, err) {
		return
	}

	httputil.RespondOK(ctx, AuthPostResponse{
		Tokens: AuthTokens{
			Authorization: jwt,
		},
		User: *user,
	})
}
