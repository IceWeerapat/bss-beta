package staff

import (
	"context"

	"fx.prodigy9.co/data"
)

type DeleteStaffForm struct {
	ID int `in:"path=id"`
}

func (d *DeleteStaffForm) DeleteChangeDetail(ctx context.Context) (err error) {
	return data.Run(ctx, func(s data.Scope) error {
		sql := `
			DELETE FROM staffs
			WHERE id = $1
		`
		return data.Exec(ctx, sql, d.ID)
	})
}
