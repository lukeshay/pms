package repositories

import (
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/lukeshay/pms/pkg/models"
)

type BookRepository struct {
	db *sqlx.DB
}

func NewBookRepository(db *sqlx.DB) *BookRepository {
	return &BookRepository{
		db: db,
	}
}

func (r *BookRepository) Get(userId, id string) (*models.Book, error) {
	book := models.Book{}
	err := r.db.Get(&book, "SELECT * FROM books WHERE user_id = $1 AND id = $2", userId, id)

	return &book, err
}

func (r *BookRepository) ListByUserId(userId string) (*[]models.Book, error) {
	books := []models.Book{}
	err := r.db.Select(&books, "SELECT * FROM books WHERE user_id = $1", userId)

	return &books, err
}

func (r *BookRepository) Insert(book *models.Book) error {
	book.CreatedAt = time.Now()
	book.UpdatedAt = time.Now()

	query := `
		INSERT INTO books (id, user_id, title, author, rating, purchased_at, finished_at, created_at, created_by, updated_at, updated_by) 
		VALUES (:id, :user_id, :title, :author, :rating, :purchased_at, :finished_at, :created_at, :created_by, :updated_at, :updated_by)
	`

	_, err := r.db.NamedExec(query, book)

	return err
}

func (r *BookRepository) Update(book *models.Book) error {
	book.UpdatedAt = time.Now()

	query := `
		UPDATE books 
		SET title = :title, author = :author, rating = :rating, purchased_at = :purchased_at, finished_at = :finished_at, created_at = :created_at, created_by = :created_by, updated_at = :updated_at, updated_by = :updated_by
		WHERE id = :id AND user_id = :user_id
	`

	_, err := r.db.NamedExec(query, book)

	return err
}

func (r *BookRepository) Delete(userId, id string) error {
	_, err := r.db.Exec("DELETE FROM books WHERE user_id = $1 AND id = $2", userId, id)

	return err
}
