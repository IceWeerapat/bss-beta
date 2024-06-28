package admin

import (
	"time"

	"fx.prodigy9.co/app"
)

var App = app.Build().
	Controllers(AdminCtr{})

type Admin struct {
	ID        int       `json:"id" db:"id"`
	Username  string    `json:"username" db:"username"`
	Password  string    `json:"-" db:"password"`
	FirstName string    `json:"first_name" db:"first_name"`
	LastName  string    `json:"last_name" db:"last_name"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
