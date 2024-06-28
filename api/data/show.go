package data

import (
	"context"

	d "fx.prodigy9.co/data"
)

type ShowDatasForm struct {
	ID string `in:"path=license_number"`
}

func (s *ShowDatasForm) ShowData(ctx context.Context, data *Data) (err error) {
	query := `
		SELECT * 
		FROM datas 
		WHERE license_number=$1
	`
	return d.Get(ctx, data, query, s.ID)
}
