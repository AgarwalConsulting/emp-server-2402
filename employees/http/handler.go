package http

import (
	"github.com/gorilla/mux"

	"algogrit.com/emp-server/employees/service"
)

type Handler struct {
	v1 service.EmployeeService
	*mux.Router
	// Router *mux.Router
}

func (h *Handler) SetupRoutes(r *mux.Router) {
	r.HandleFunc("/v1/employees", h.CreateV1).Methods("POST")
	r.HandleFunc("/v1/employees", h.IndexV1).Methods("GET")

	h.Router = r
}

func NewHandler(v1 service.EmployeeService) Handler {
	h := Handler{v1: v1}

	r := mux.NewRouter()

	h.SetupRoutes(r)

	return h
}
