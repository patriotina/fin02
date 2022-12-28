package service

import (
	"fin02/internal/action"
	"fin02/pkg/response"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func New() ServerImplementation {
	return ServerImplementation{}
}

type ServerImplementation struct {
}

func (s ServerImplementation) BuildRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/hello_world", func(w http.ResponseWriter, r *http.Request) {
		result, err := action.GetHelloWorld()
		response.Make(w, result, err)
	})
	r.Get("/task/Циклическая ротация", func(w http.ResponseWriter, r *http.Request) {
		result, err := action.GetTask1()
		response.Make(w, result, err)
	})
	//https://kuvaev-ituniversity.vps.elewise.com/tasks/%D0%A6%D0%B8%D0%BA%D0%BB%D0%B8%D1%87%D0%B5%D1%81%D0%BA%D0%B0%D1%8F%20%D1%80%D0%BE%D1%82%D0%B0%D1%86%D0%B8%D1%8F
	return r
}
