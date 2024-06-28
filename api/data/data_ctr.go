package data

import (
	"net/http"

	"fx.prodigy9.co/config"
	"fx.prodigy9.co/httpserver/controllers"
	"fx.prodigy9.co/httpserver/middlewares"
	"fx.prodigy9.co/httpserver/render"
	"github.com/ggicci/httpin"
	"github.com/go-chi/chi/v5"
)

type DatasCtr struct{}

var _ controllers.Interface = DatasCtr{}

func (d DatasCtr) Mount(cfg *config.Source, r chi.Router) error {
	r.Route("/api/data", func(sr chi.Router) {
		sr.With(middlewares.Validation(CreateDatasForm{})...).Post("/", d.Create)
		sr.With(middlewares.Validation(UpdateDatasForm{})...).Patch("/{id}", d.Update)
		sr.With(middlewares.Validation(DeleteDatasForm{})...).Delete("/{id}", d.Delete)
		sr.With(middlewares.Validation(ShowDatasForm{})...).Get("/{license_number}", d.Show)
		sr.With(middlewares.Validation(ListDatasForm{})...).Get("/list", d.List)

		sr.With(middlewares.Validation(ListDatasByCriteriaForm{})...).Get("/list-criteria", d.ListDataByCriteria)

		sr.With(middlewares.Validation(SummaryDataForm{})...).Get("/summary/{staff_id}", d.Summary)
		sr.With(middlewares.Validation(SummaryDataForm{})...).Get("/summary", d.Summary)
	})
	return nil
}

func (d DatasCtr) Create(resp http.ResponseWriter, req *http.Request) {
	create := req.Context().Value(httpin.Input).(*CreateDatasForm)
	data := &Data{}
	if err := create.CreateData(req.Context(), &resp, data); err != nil {
		render.Error(resp, req, 400, err)
	} else {
		render.JSON(resp, req, data)
	}
}

func (d DatasCtr) Update(resp http.ResponseWriter, req *http.Request) {
	update := req.Context().Value(httpin.Input).(*UpdateDatasForm)
	data := &Data{}
	if err := update.UpdateData(req.Context(), &resp, data); err != nil {
		render.Error(resp, req, 400, err)
	} else {
		render.JSON(resp, req, data)
	}
}

func (d DatasCtr) Delete(resp http.ResponseWriter, req *http.Request) {
	delete := req.Context().Value(httpin.Input).(*DeleteDatasForm)
	if err := delete.DeleteData(req.Context()); err != nil {
		render.Error(resp, req, http.StatusInternalServerError, err)
		return
	}
	render.JSON(resp, req, true)
}

func (d DatasCtr) Show(resp http.ResponseWriter, req *http.Request) {
	show := req.Context().Value(httpin.Input).(*ShowDatasForm)
	data := &Data{}
	if err := show.ShowData(req.Context(), data); err != nil {
		render.Error(resp, req, 400, err)
	} else {
		render.JSON(resp, req, data)
	}
}

func (d DatasCtr) List(resp http.ResponseWriter, req *http.Request) {
	list := req.Context().Value(httpin.Input).(*ListDatasForm)
	datas := &[]*Data{}
	if err := list.ListData(req.Context(), datas); err != nil {
		render.Error(resp, req, 400, err)
	} else {
		render.JSON(resp, req, datas)
	}
}

func (d DatasCtr) ListDataByCriteria(resp http.ResponseWriter, req *http.Request) {
	form := req.Context().Value(httpin.Input).(*ListDatasByCriteriaForm)
	datas := &[]*DataWithStaffResponse{}
	page := &render.Pagination{CurrentPage: form.Page, LimitPerPage: form.Limit}
	if err := form.ListDataByCriteria(req.Context(), datas, page); err != nil {
		render.Error(resp, req, http.StatusInternalServerError, err)
		return
	} else {
		render.JSON(resp, req, datas, page)
	}
}

func (d DatasCtr) Summary(resp http.ResponseWriter, req *http.Request) {
	summary := req.Context().Value(httpin.Input).(*SummaryDataForm)
	res := &SummaryResponse{}
	if err := summary.Summary(req.Context(), res); err != nil {
		render.Error(resp, req, 400, err)
	} else {
		render.JSON(resp, req, res)
	}
}
