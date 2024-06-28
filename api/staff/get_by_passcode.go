package staff

import (
	"context"

	"fx.prodigy9.co/data"
)

type PinCodeStaffForm struct {
	PinCode int `in:"path=pin_code"`
}

func (s *PinCodeStaffForm) ShowStaffByPinCode(ctx context.Context, staff *Staff) (err error) {
	query := `
		SELECT * 
		FROM staffs 
		WHERE pin_code=$1
	`
	return data.Get(ctx, staff, query, s.PinCode)
}
