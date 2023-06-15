package controllers

import (
	"github.com/lukeshay/pms/models"

	"github.com/gofiber/fiber/v2"
)

type BooksV1Controller struct {
	Controller
	basePath string
}

func NewBooksV1Controller() Controller {
	return &BooksV1Controller{
		basePath: "/books",
	}
}

func (c *BooksV1Controller) BasePath() string {
	return c.basePath
}

type BooksListData struct {
	Books []models.Book `json:"books" binding:"required"`
}

// BooksList godoc
// @Summary      List books
// @Description  get books
// @Tags         books
// @Accept       json
// @Produce      json
// @Success      200  {object}  BooksListData
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /v1/books [get]
func (c *BooksV1Controller) List(ctx *fiber.Ctx) error {
	data := BooksListData{
		Books: []models.Book{
			{
				Id:     "1",
				Title:  "Never Finished",
				Author: "David Goggins",
				Rating: 5,
			},
			{
				Id:     "2",
				Title:  "Can't Hurt Me",
				Author: "David Goggins",
				Rating: 4,
			},
		},
	}

	return ctx.JSON(data)
}

type BooksGetData struct {
	Book models.Book `json:"book" binding:"required"`
}

// BooksGet godoc
// @Summary      Get book by id
// @Description  get book by id
// @Tags         books
// @Accept       json
// @Produce      json
// @Param 			 id path string true "Book ID"
// @Success      200  {object}  BooksGetData
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /v1/book/{id} [get]
func (c *BooksV1Controller) Get(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusNotImplemented).JSON(map[string]string{
		"message": "Not implemented",
	})
}

// BooksPost godoc
// @Summary      Create book
// @Description  post book
// @Tags         books
// @Accept       json
// @Produce      json
// @Success      200  {object}  BooksGetData
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /v1/book [post]
func (c *BooksV1Controller) Post(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusNotImplemented).JSON(map[string]string{
		"message": "Not implemented",
	})
}

// BooksPut godoc
// @Summary      Update book by id
// @Description  put book by id
// @Tags         books
// @Accept       json
// @Produce      json
// @Param 			 id path string true "Book ID"
// @Success      200  {object}  BooksGetData
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /v1/book/{id} [put]
func (c *BooksV1Controller) Put(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusNotImplemented).JSON(map[string]string{
		"message": "Not implemented",
	})
}

// BooksDelete godoc
// @Summary      Update book by id
// @Description  delete book by id
// @Tags         books
// @Accept       json
// @Produce      json
// @Param 			 id path string true "Book ID"
// @Success      200  {object}  BooksGetData
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /v1/book/{id} [delete]
func (c *BooksV1Controller) Delete(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusNotImplemented).JSON(map[string]string{
		"message": "Not implemented",
	})
}
