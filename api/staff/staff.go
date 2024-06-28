package staff

import (
	"time"

	"fx.prodigy9.co/app"
)

var App = app.Build().
	Controllers(StaffCtr{})

type Staff struct {
	ID        int       `json:"id" db:"id"`
	PinCode   int       `json:"pin_code" db:"pin_code"`
	FirstName string    `json:"first_name" db:"first_name"`
	LastName  string    `json:"last_name" db:"last_name"`
	CreatedBy string    `json:"created_by" db:"created_by"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedBy string    `json:"updated_by" db:"updated_by"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
