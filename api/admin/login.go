package admin

import (
	"context"
	"net/http"

	"fx.prodigy9.co/data"
)

type LoginRequest struct {
	Payload struct {
		Username string `json:"username"`
		Password string `json:"password"`
	} `in:"body=json"`
}

func (l *LoginRequest) Login(ctx context.Context, resp *http.ResponseWriter, admin *Admin) (err error) {
	query := `
	SELECT *
	FROM admins AS a
	WHERE a.username = $1 AND a.password = $2
	`
	return data.Get(ctx, admin, query, l.Payload.Username, l.Payload.Password)
}
