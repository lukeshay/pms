package httputil

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HTTPErrorError struct {
	Message string `json:"message"`
}

type HTTPError struct {
	Error HTTPErrorError `json:"error" binding:"required"`
}

func ShouldBindJSON(ctx *gin.Context, obj interface{}) bool {
	err := ctx.ShouldBindJSON(obj)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, HTTPError{
			Error: HTTPErrorError{
				Message: err.Error(),
			},
		})
		return false
	}

	return true
}

func RespondOK[T any](ctx *gin.Context, body T) {
	ctx.JSON(http.StatusOK, body)
}

func RespondNotFound(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusNotFound, HTTPError{
		Error: HTTPErrorError{
			Message: msg,
		},
	})
}

func RespondInternalServerError(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusInternalServerError, HTTPError{
		Error: HTTPErrorError{
			Message: msg,
		},
	})
}

func RespondNotFoundQuery[T any](ctx *gin.Context, msg string, obj *T, err error) bool {
	if err != nil {
		log.Fatalln(err.Error())

		RespondInternalServerError(ctx, err.Error())

		return false
	}

	if obj == nil {
		RespondNotFound(ctx, msg)

		return false
	}

	return true
}

func RespondInternalServerErrorQuery[T any](ctx *gin.Context, msg string, obj *T, err error) bool {
	if err != nil || obj == nil {
		log.Fatalln(err.Error())

		RespondInternalServerError(ctx, msg)

		return false
	}

	return true
}

func RespondForbidden(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusForbidden, HTTPError{
		Error: HTTPErrorError{
			Message: msg,
		},
	})
}
