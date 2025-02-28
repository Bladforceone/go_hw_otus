package repository

import (
	"context"

	"github.com/Bladforceone/go_hw_otus/hw15_go_sql/internal/db"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	Queries *db.Queries
	DB      *pgxpool.Pool
}

func NewRepository(dbPool *pgxpool.Pool) *Repository {
	return &Repository{
		Queries: db.New(dbPool),
		DB:      dbPool,
	}
}

func (r *Repository) ProductCreate(ctx context.Context, name string,
	price pgtype.Numeric, stock int32,
) (*db.Product, error) {
	product, err := r.Queries.ProductCreate(ctx, db.ProductCreateParams{
		Name:  name,
		Price: price,
		Stock: stock,
	})
	return product, err
}

func (r *Repository) ProductGet(ctx context.Context, id int32) (*db.Product, error) {
	return r.Queries.ProductGet(ctx, id)
}

func (r *Repository) ProductDelete(ctx context.Context, id int32) error {
	return r.Queries.ProductDelete(ctx, id)
}
