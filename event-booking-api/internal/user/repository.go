package user

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Create(ctx context.Context, user *User) error {
	query := `
        INSERT INTO users (name, email, password_hash)
        VALUES ($1, $2, $3)
        RETURNING id, created_at
    `

	err := r.db.QueryRow(
		ctx,
		query,
		user.Name,
		user.Email,
		user.PasswordHash,
	).Scan(&user.ID, &user.CreatedAt)

	if err != nil {
		return fmt.Errorf("create user: %w", err)
	}

	return nil
}
