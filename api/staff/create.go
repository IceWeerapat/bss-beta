package staff

import (
	"context"
	"net/http"

	"fx.prodigy9.co/data"
)

type CreateStaffForm struct {
	Payload struct {
		PinCode   int    `json:"pin_code"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		CreatedBy string `json:"created_by"`
		UpdatedBy string `json:"updated_by"`
	} `in:"body=json"`
}

func (c *CreateStaffForm) CreateStaff(ctx context.Context, resp *http.ResponseWriter, staff *Staff) (err error) {
	query := `
		INSERT INTO staffs (
			pin_code, 
			first_name, 
			last_name, 
			created_by, 
			updated_by
		)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING *
	`
	return data.Get(ctx, staff, query, c.Payload.PinCode, c.Payload.FirstName,
		c.Payload.LastName, c.Payload.CreatedBy, c.Payload.UpdatedBy)
}
