package staff

import (
	"context"

	"fx.prodigy9.co/data"
)

type ShowStaffForm struct {
	ID int `in:"path=id"`
}

func (s *ShowStaffForm) ShowStaff(ctx context.Context, staff *Staff) (err error) {
	query := `
		SELECT * 
		FROM staffs 
		WHERE id=$1
	`
	return data.Get(ctx, staff, query, s.ID)
}
