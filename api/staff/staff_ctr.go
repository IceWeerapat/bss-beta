package staff

import (
	"net/http"

	"fx.prodigy9.co/config"
	"fx.prodigy9.co/httpserver/controllers"
	"fx.prodigy9.co/httpserver/middlewares"
	"fx.prodigy9.co/httpserver/render"
	"github.com/ggicci/httpin"
	"github.com/go-chi/chi/v5"
)

type StaffCtr struct{}

var _ controllers.Interface = StaffCtr{}

func (s StaffCtr) Mount(cfg *config.Source, r chi.Router) error {
	r.Route("/api/staff", func(sr chi.Router) {
		sr.With(middlewares.Validation(CreateStaffForm{})...).Post("/", s.Create)
		sr.With(middlewares.Validation(UpdateStaffForm{})...).Patch("/{id}", s.Update)
		sr.With(middlewares.Validation(DeleteStaffForm{})...).Delete("/{id}", s.Delete)
		sr.With(middlewares.Validation(ShowStaffForm{})...).Get("/{id}", s.Show)
		sr.With(middlewares.Validation(ListStaffForm{})...).Post("/list", s.List)
		sr.With(middlewares.Validation(ListStaffForm{})...).Post("/list-order-name", s.ListOrderName)
		sr.With(middlewares.Validation(PinCodeStaffForm{})...).Get("/pin-code/{pin_code}", s.ShowByPinCode)
	})
	return nil
}

func (s StaffCtr) Create(resp http.ResponseWriter, req *http.Request) {
	create := req.Context().Value(httpin.Input).(*CreateStaffForm)
	staff := &Staff{}
	if err := create.CreateStaff(req.Context(), &resp, staff); err != nil {
		render.Error(resp, req, 400, err)
	} else {
		render.JSON(resp, req, staff)
	}
}

func (s StaffCtr) Update(resp http.ResponseWriter, req *http.Request) {
	update := req.Context().Value(httpin.Input).(*UpdateStaffForm)
	staff := &Staff{}
	if err := update.UpdateStaff(req.Context(), &resp, staff); err != nil {
		render.Error(resp, req, 400, err)
	} else {
		render.JSON(resp, req, staff)
	}
}

func (s StaffCtr) Delete(resp http.ResponseWriter, req *http.Request) {
	delete := req.Context().Value(httpin.Input).(*DeleteStaffForm)
	if err := delete.DeleteChangeDetail(req.Context()); err != nil {
		render.Error(resp, req, http.StatusInternalServerError, err)
		return
	}
	render.JSON(resp, req, true)
}

func (s StaffCtr) Show(resp http.ResponseWriter, req *http.Request) {
	show := req.Context().Value(httpin.Input).(*ShowStaffForm)
	staff := &Staff{}
	if err := show.ShowStaff(req.Context(), staff); err != nil {
		render.Error(resp, req, 400, err)
	} else {
		render.JSON(resp, req, staff)
	}
}

func (s StaffCtr) List(resp http.ResponseWriter, req *http.Request) {
	list := req.Context().Value(httpin.Input).(*ListStaffForm)
	staffs := &[]*Staff{}
	if err := list.ListStaff(req.Context(), staffs); err != nil {
		render.Error(resp, req, 400, err)
	} else {
		render.JSON(resp, req, staffs)
	}
}

func (s StaffCtr) ListOrderName(resp http.ResponseWriter, req *http.Request) {
	list := req.Context().Value(httpin.Input).(*ListStaffForm)
	staffs := &[]*Staff{}
	if err := list.ListStaffOrderByName(req.Context(), staffs); err != nil {
		render.Error(resp, req, 400, err)
	} else {
		render.JSON(resp, req, staffs)
	}
}

func (s StaffCtr) ShowByPinCode(resp http.ResponseWriter, req *http.Request) {
	show := req.Context().Value(httpin.Input).(*PinCodeStaffForm)
	staff := &Staff{}
	if err := show.ShowStaffByPinCode(req.Context(), staff); err != nil {
		render.Error(resp, req, 400, err)
	} else {
		render.JSON(resp, req, staff)
	}
}
