package staff

import (
	"context"
	"net/http"

	"fx.prodigy9.co/data"
)

type UpdateStaffForm struct {
	ID      int `in:"path=id"`
	Payload struct {
		PinCode   int    `json:"pin_code"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		UpdatedBy string `json:"updated_by"`
	} `in:"body=json"`
}

func (u *UpdateStaffForm) UpdateStaff(ctx context.Context, resp *http.ResponseWriter, staff *Staff) error {
	query := `
		UPDATE staffs
		SET pin_code=$1, first_name=$2, last_name=$3, updated_by=$4
		WHERE id=$5
		RETURNING *
	`
	return data.Get(ctx, staff, query, u.Payload.PinCode, u.Payload.FirstName, u.Payload.LastName, u.Payload.UpdatedBy, u.ID)
}
