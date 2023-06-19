package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lukeshay/pms/pkg/auth"
	"github.com/lukeshay/pms/pkg/httputil"
	"github.com/lukeshay/pms/pkg/models"
	"github.com/lukeshay/pms/pkg/repositories"
)

type BooksV1Controller struct {
	ModelController
	basePath       string
	bookRepository *repositories.BookRepository
	auth           *auth.Auth
}

type BooksV1ControllerInput struct {
	BookRepository *repositories.BookRepository
}

func NewBooksV1Controller(input BooksV1ControllerInput) ModelController {
	return &BooksV1Controller{
		basePath:       "/books",
		bookRepository: input.BookRepository,
	}
}

func (c *BooksV1Controller) BasePath() string {
	return c.basePath
}

type BooksListResponse struct {
	Books []models.Book `json:"books" binding:"required"`
}

// BooksList godoc
// @Summary      List books
// @Description  get books
// @Tags         books
// @Accept       json
// @Produce      json
// @Success      200  {object}  BooksListResponse
// @Failure      400  {object}  httputil.HTTPError
// @Failure      403  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /v1/books/ [get]
func (c *BooksV1Controller) List(ctx *gin.Context) {
	books, err := c.bookRepository.ListByUserId(auth.RequireClaims(ctx).User.Id)
	if !httputil.RespondInternalServerErrorQuery(ctx, "Unknown error", books, err) {
		return
	}

	httputil.RespondOK(ctx, BooksListResponse{
		Books: *books,
	})
}

type BooksGetResponse struct {
	Book models.Book `json:"book" binding:"required"`
}

// BooksGet godoc
// @Summary      Get book by id
// @Description  get book by id
// @Tags         books
// @Accept       json
// @Produce      json
// @Param 			 id path string true "Book ID"
// @Success      200  {object}  BooksGetResponse
// @Failure      400  {object}  httputil.HTTPError
// @Failure      403  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /v1/books/{id}/ [get]
func (c *BooksV1Controller) Get(ctx *gin.Context) {
	bookId := ctx.Param("id")
	userId := auth.RequireClaims(ctx).User.Id

	book, err := c.bookRepository.Get(userId, bookId)
	if !httputil.RespondNotFoundQuery(ctx, fmt.Sprintf("Book with id %s not found", bookId), book, err) {
		return
	}

	ctx.JSON(http.StatusOK, BooksGetResponse{
		Book: *book,
	})
}

type BooksPostRequest struct {
	UserId      string     `json:"userId"`
	Title       string     `json:"title" binding:"required"`
	Author      string     `json:"author" binding:"required"`
	Rating      int8       `json:"rating"`
	PurchasedAt *time.Time `json:"purchasedAt"`
	FinishedAt  *time.Time `json:"finishedAt"`
}

// BooksPost godoc
// @Summary      Create book
// @Description  post book
// @Tags         books
// @Accept       json
// @Produce      json
// @Param			   book	body		  BooksPostRequest    true  "Create book"
// @Success      200  {object}  BooksGetResponse
// @Failure      400  {object}  httputil.HTTPError
// @Failure      403  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /v1/books/ [post]
func (c *BooksV1Controller) Post(ctx *gin.Context) {
	var request BooksPostRequest
	if !httputil.ShouldBindJSON(ctx, &request) {
		return
	}

	userId := auth.RequireClaims(ctx).User.Id
	bookId := uuid.NewString()

	book := &models.Book{
		Id:          bookId,
		UserId:      userId,
		Title:       request.Title,
		Author:      request.Author,
		Rating:      request.Rating,
		PurchasedAt: request.PurchasedAt,
		FinishedAt:  request.FinishedAt,
		Timestamps: models.Timestamps{
			CreatedBy: userId,
			UpdatedBy: userId,
		},
	}

	err := c.bookRepository.Insert(book)
	if !httputil.RespondInternalServerErrorQuery(ctx, "Unknown error", book, err) {
		return
	}

	ctx.JSON(http.StatusCreated, BooksGetResponse{
		Book: *book,
	})
}

// BooksPut godoc
// @Summary      Update book by id
// @Description  put book by id
// @Tags         books
// @Accept       json
// @Produce      json
// @Param 			 id path string true "Book ID"
// @Param			   book	body		  BooksPostRequest    true  "Update book"
// @Success      200  {object}  BooksGetResponse
// @Failure      400  {object}  httputil.HTTPError
// @Failure      403  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /v1/books/{id}/ [put]
func (c *BooksV1Controller) Put(ctx *gin.Context) {
	var request BooksPostRequest
	if !httputil.ShouldBindJSON(ctx, &request) {
		return
	}

	userId := auth.RequireClaims(ctx).User.Id
	bookId := ctx.Param("id")

	book := &models.Book{
		Id:          bookId,
		UserId:      userId,
		Title:       request.Title,
		Author:      request.Author,
		Rating:      request.Rating,
		PurchasedAt: request.PurchasedAt,
		FinishedAt:  request.FinishedAt,
		Timestamps: models.Timestamps{
			CreatedBy: userId,
			UpdatedBy: userId,
		},
	}

	err := c.bookRepository.Update(book)
	if !httputil.RespondInternalServerErrorQuery(ctx, "Unknown error", book, err) {
		return
	}

	ctx.JSON(http.StatusOK, BooksGetResponse{
		Book: *book,
	})
}

type BooksDeleteResponse struct {
	BookId string `json:"bookId" binding:"required"`
}

// BooksDelete godoc
// @Summary      Delete book by id
// @Description  delete book by id
// @Tags         books
// @Accept       json
// @Produce      json
// @Param 			 id path string true "Book ID"
// @Success      200  {object}  BooksDeleteResponse
// @Failure      500  {object}  httputil.HTTPError
// @Router       /v1/books/{id}/ [delete]
func (c *BooksV1Controller) Delete(ctx *gin.Context) {
	userId := auth.RequireClaims(ctx).User.Id
	bookId := ctx.Param("id")

	if err := c.bookRepository.Delete(userId, bookId); err != nil {
		httputil.RespondInternalServerError(ctx, "Unknown error")
		return
	}

	httputil.RespondOK(ctx, BooksDeleteResponse{
		BookId: bookId,
	})
}
