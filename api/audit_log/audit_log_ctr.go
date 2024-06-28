package auditlog

import (
	"net/http"

	"fx.prodigy9.co/config"
	"fx.prodigy9.co/httpserver/controllers"
	"fx.prodigy9.co/httpserver/middlewares"
	"fx.prodigy9.co/httpserver/render"
	"github.com/ggicci/httpin"
	"github.com/go-chi/chi/v5"
)

type AuditLogCtr struct{}

var _ controllers.Interface = AuditLogCtr{}

func (a AuditLogCtr) Mount(cfg *config.Source, r chi.Router) error {
	r.Route("/api/audit-log", func(sr chi.Router) {
		sr.With(middlewares.Validation(ListAuditLogForm{})...).Post("/list", a.List)
	})
	return nil
}

func (a AuditLogCtr) List(resp http.ResponseWriter, req *http.Request) {
	list := req.Context().Value(httpin.Input).(*ListAuditLogForm)
	auditLogs := &[]*AuditLogResult{}
	if err := list.ListAuditLogs(req.Context(), auditLogs); err != nil {
		render.Error(resp, req, 400, err)
	} else {
		render.JSON(resp, req, auditLogs)
	}
}
