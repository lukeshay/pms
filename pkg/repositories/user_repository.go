package repositories

import (
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/lukeshay/pms/pkg/models"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) CreateTable() {
	_, err := r.db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id TEXT PRIMARY KEY,
		email TEXT NOT NULL UNIQUE,
		email_verified BOOLEAN NOT NULL,
		first_name TEXT NOT NULL,
		last_name TEXT NOT NULL,
		password TEXT NOT NULL,
		created_at TIMESTAMP NOT NULL,
		created_by TEXT NOT NULL,
		updated_at TIMESTAMP NOT NULL,
		updated_by TEXT NOT NULL,
		deleted_at TIMESTAMP,
		deleted_by TEXT
	)`)

	if err != nil {
		panic(err)
	}
}

func (r *UserRepository) Get(id string) (*models.User, error) {
	user := models.User{}
	err := r.db.Get(&user, "SELECT * FROM users WHERE id = $1", id)

	return &user, err
}

func (r *UserRepository) GetByEmail(email string) (*models.User, error) {
	user := models.User{}
	err := r.db.Get(&user, "SELECT * FROM users WHERE email = $1", email)

	return &user, err
}

func (r *UserRepository) Insert(user *models.User) error {
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	query := `
		INSERT INTO users (id, email, email_verified, first_name, last_name, password, created_at, created_by, updated_at, updated_by) 
		VALUES (:id, :email, :email_verified, :first_name, :last_name, :password, :created_at, :created_by, :updated_at, :updated_by)
	`

	_, err := r.db.NamedExec(query, user)

	return err
}

func (r *UserRepository) Update(user *models.User) error {
	user.UpdatedAt = time.Now()

	query := `
		UPDATE users 
		SET email = :email, email_verified = :email_verified, first_name = :first_name, last_name = :last_name, password = :password, updated_at = :updated_at, updated_by = :updated_by
		WHERE id = :id
	`

	_, err := r.db.NamedExec(query, user)

	return err
}
