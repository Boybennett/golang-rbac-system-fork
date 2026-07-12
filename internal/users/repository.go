package users

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	// Import the auto-generated sqlc code locally
	usersdb "github.com/Steve-s-Circle-on-System-Design/golang-rbac-system/internal/users/sqlc"
)

// Repository encapsulates all database storage access for the users feature.
type Repository struct {
	pool    *pgxpool.Pool
	queries *usersdb.Queries
}

// NewRepository is exactly where you initialize your sqlc queries.
func NewRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{
		pool:    pool,
		queries: usersdb.New(pool), // <-- sqlc initialization happens right here
	}
}

// Example method utilizing the initialized sqlc queries
func (r *Repository) FindByID(ctx context.Context, email string) (usersdb.User, error) {
	return r.queries.GetUserByEmail(ctx, email)
}
