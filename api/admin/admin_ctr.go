package admin

import (
	"net/http"

	"fx.prodigy9.co/config"
	"fx.prodigy9.co/httpserver/controllers"
	"fx.prodigy9.co/httpserver/middlewares"
	"fx.prodigy9.co/httpserver/render"
	"github.com/ggicci/httpin"
	"github.com/go-chi/chi/v5"
)

type AdminCtr struct{}

var _ controllers.Interface = AdminCtr{}

func (a AdminCtr) Mount(cfg *config.Source, r chi.Router) error {
	r.Route("/api/admin/login", func(sr chi.Router) {
		sr.With(middlewares.Validation(LoginRequest{})...).Post("/", a.Login)

	})
	return nil
}

func (a AdminCtr) Login(resp http.ResponseWriter, req *http.Request) {
	login := req.Context().Value(httpin.Input).(*LoginRequest)
	admin := &Admin{}
	if err := login.Login(req.Context(), &resp, admin); err != nil {
		render.Error(resp, req, 400, err)
	} else {
		render.JSON(resp, req, admin)
	}
}
