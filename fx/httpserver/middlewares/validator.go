package middlewares

import (
	"errors"
	"net/http"
	"reflect"

	"fx.prodigy9.co/config"
	"fx.prodigy9.co/httpserver/httperrors"
	"fx.prodigy9.co/httpserver/render"
	"github.com/ggicci/httpin"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

func Validator(cfg *config.Source) func(http.Handler) http.Handler {
	httpin.UseGochiURLParam("path", chi.URLParam)
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(resp http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(resp, r)
		})
	}
}

func Validation(form interface{}) []func(next http.Handler) http.Handler {
	return []func(next http.Handler) http.Handler{
		requestToStruct(form),
		structValidator(form),
	}
}

func errorHandler(rw http.ResponseWriter, r *http.Request, err error) {
	if httpinErr, ok := errors.Unwrap(err).(*httpin.InvalidFieldError); ok {
		unwraped := httpinErr.Unwrap()
		if unwraped.Error() == "EOF" {
			unwraped = errors.New("error parsing body")
		}
		render.Error(rw, r, 400, unwraped)
	}
}

// requestToStruct reads the request data into the struct
func requestToStruct(form interface{}) func(http.Handler) http.Handler {
	return httpin.NewInput(form, httpin.WithErrorHandler(errorHandler))
}

// structValidator validates the root struct with go validator
func structValidator(form interface{}) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			val := validator.New()
			input := req.Context().Value(httpin.Input)

			inputValue := reflect.ValueOf(input)
			inputStruct := inputValue.Elem().Interface()

			err := val.Struct(inputStruct)

			if err != nil {
				render.Error(rw, req, 400, httperrors.NewError(httperrors.ErrBadRequest.Code(), err.Error()))
				return
			}
			next.ServeHTTP(rw, req)
		})
	}
}
