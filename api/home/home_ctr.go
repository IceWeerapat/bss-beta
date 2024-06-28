package home

import (
	"net/http"

	"fx.prodigy9.co/config"
	"fx.prodigy9.co/httpserver/controllers"
	"fx.prodigy9.co/httpserver/render"
	"github.com/go-chi/chi/v5"
)

type HomeCtr struct{}

var _ controllers.Interface = HomeCtr{}

// Mount implements controllers.Interface
func (h HomeCtr) Mount(cfg *config.Source, router chi.Router) error {
	router.Get("/", h.Index)
	return nil
}

func (h HomeCtr) Index(resp http.ResponseWriter, req *http.Request) {
	render.JSON(resp, req, NewStatus())
}
