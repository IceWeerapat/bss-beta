package data

import (
	"context"

	"fx.prodigy9.co/data"
)

type DeleteDatasForm struct {
	ID int `in:"path=id"`
}

func (d *DeleteDatasForm) DeleteData(ctx context.Context) error {
	return data.Run(ctx, func(s data.Scope) error {
		query := `
		DELETE FROM datas
		WHERE id=$1
	`
		return data.Exec(ctx, query, d.ID)
	})
}
