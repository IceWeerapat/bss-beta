package staff

import (
	"context"

	"fx.prodigy9.co/data"
)

type ListStaffForm struct {
}

func (l *ListStaffForm) ListStaff(ctx context.Context, staffs *[]*Staff) (err error) {
	query := `
		SELECT * 
		FROM staffs 
		ORDER BY ID ASC
	`
	return data.Select(ctx, staffs, query)
}

func (l *ListStaffForm) ListStaffOrderByName(ctx context.Context, staffs *[]*Staff) (err error) {
	query := `
		SELECT * 
		FROM staffs 
		ORDER BY first_name , last_name ASC
	`
	return data.Select(ctx, staffs, query)
}
