package controllers

import (
	"github.com/gofiber/fiber/v2"
)

type Controller interface {
	// Returns the base path for this controller
	//  GET /<BasePath>/
	//  Responds: JSON, HTML
	List(ctx *fiber.Ctx) error
	// Returns the requested item
	//  GET /<BasePath>/{id}/
	//  Responds: JSON, HTML
	Get(ctx *fiber.Ctx) error
	// Creates a new item
	//  POST /<BasePath>/
	//  Accepts: JSON, FORM
	//  Responds: JSON
	Post(ctx *fiber.Ctx) error
	// Updates an item
	//  PUT /<BasePath>/{id}/
	//  Accepts: JSON, FORM
	//  Responds: JSON
	Put(ctx *fiber.Ctx) error

	// Deletes an item
	//  DELETE /<BasePath>/{id}/
	//  Responds: JSON, REDIRECT
	Delete(ctx *fiber.Ctx) error

	BasePath() string
}
