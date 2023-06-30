package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lukeshay/pms/pkg/auth"
	"github.com/lukeshay/pms/pkg/httputil"
	"github.com/lukeshay/pms/pkg/models"
	"github.com/lukeshay/pms/pkg/repositories"
)

type UsersV1Controller struct {
	basePath       string
	userRepository *repositories.UserRepository
}

type UsersV1ControllerInput struct {
	UserRepository *repositories.UserRepository
}

func NewUsersV1Controller(input UsersV1ControllerInput) *UsersV1Controller {
	return &UsersV1Controller{
		basePath:       "/users",
		userRepository: input.UserRepository,
	}
}

func (c *UsersV1Controller) BasePath() string {
	return c.basePath
}

type UsersGetResponse struct {
	User models.User `json:"user" binding:"required"`
}

// UsersGet godoc
// @Summary      Get user by id
// @Description  get user by id
// @Tags         users
// @Accept       json
// @Produce      json
// @Param 			 id path string true "User ID"
// @Success      200  {object}  UsersGetResponse
// @Failure      400  {object}  httputil.HTTPError
// @Failure      403  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /v1/users/{id}/ [get]
func (c *UsersV1Controller) Get(ctx *gin.Context) {
	userId := ctx.Param("id")
	currentUserId := auth.RequireClaims(ctx).User.Id

	if userId != currentUserId {
		httputil.RespondForbidden(ctx, "You do not have permission to access this user")
		return
	}

	user, err := c.userRepository.Get(userId)
	if !httputil.RespondNotFoundQuery(ctx, fmt.Sprintf("User with id %s not found", userId), user, err) {
		return
	}

	ctx.JSON(http.StatusOK, UsersGetResponse{
		User: *user,
	})
}

type UsersPostRequest struct {
	UserId    string `json:"userId"`
	Email     string `json:"email" binding:"required"`
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
}

// UsersPut godoc
// @Summary      Update user by id
// @Description  put user by id
// @Tags         users
// @Accept       json
// @Produce      json
// @Param 			 id path string true "User ID"
// @Param			   user	body		  UsersPostRequest    true  "Update user"
// @Success      200  {object}  UsersGetResponse
// @Failure      400  {object}  httputil.HTTPError
// @Failure      403  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /v1/users/{id}/ [put]
func (c *UsersV1Controller) Put(ctx *gin.Context) {
	var request UsersPostRequest
	if !httputil.ShouldBindJSON(ctx, &request) {
		return
	}

	currentUserId := auth.RequireClaims(ctx).User.Id
	userId := ctx.Param("id")

	if userId != currentUserId {
		httputil.RespondForbidden(ctx, "You do not have permission to access this user")
		return
	}

	user := &models.User{
		Id:            userId,
		Email:         request.Email,
		EmailVerified: false,
		FirstName:     request.FirstName,
		LastName:      request.LastName,
		Timestamps: models.Timestamps{
			CreatedBy: userId,
			UpdatedBy: userId,
		},
	}

	err := c.userRepository.Update(user)
	if !httputil.RespondInternalServerErrorQuery(ctx, "Unknown error", user, err) {
		return
	}

	ctx.JSON(http.StatusOK, UsersGetResponse{
		User: *user,
	})
}
