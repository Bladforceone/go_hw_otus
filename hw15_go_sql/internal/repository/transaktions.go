package repository

import (
	"context"

	"github.com/Bladforceone/go_hw_otus/hw15_go_sql/internal/db"
)

func (r *Repository) ProductUpdateStockAndDelete(ctx context.Context, id int32, newStock int32) error {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	repo := r.Queries.WithTx(tx)

	err = repo.ProductUpdateStock(ctx, db.ProductUpdateStockParams{
		Stock: newStock,
		ID:    id,
	})
	if err != nil {
		return err
	}

	if newStock == 0 {
		err = repo.ProductDelete(ctx, id)
		if err != nil {
			return err
		}
	}

	return tx.Commit(ctx)
}
