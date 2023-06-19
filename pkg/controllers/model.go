package controllers

import (
	"github.com/gin-gonic/gin"
)

type ModelController interface {
	// Returns the base path for this controller
	//  GET /<BasePath>/
	//  Responds: JSON, HTML
	List(ctx *gin.Context)
	// Returns the requested item
	//  GET /<BasePath>/{id}/
	//  Responds: JSON, HTML
	Get(ctx *gin.Context)
	// Creates a new item
	//  POST /<BasePath>/
	//  Accepts: JSON, FORM
	//  Responds: JSON
	Post(ctx *gin.Context)
	// Updates an item
	//  PUT /<BasePath>/{id}/
	//  Accepts: JSON, FORM
	//  Responds: JSON
	Put(ctx *gin.Context)

	// Deletes an item
	//  DELETE /<BasePath>/{id}/
	//  Responds: JSON, REDIRECT
	Delete(ctx *gin.Context)

	BasePath() string
}
