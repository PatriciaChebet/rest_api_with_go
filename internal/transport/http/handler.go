package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//Handler - stores pointer to our comment service
type Handler struct {
	Router *mux.Router
}

//NewHandler - returns pointer to a Handler
func NewHandler() *Handler {
	return &Handler{}
}

//SetUpRoutes - sets up all the routes for our application
func (h *Handler) SetUpRoutes() {
	fmt.Println("Setting up Routes")
	h.Router = mux.NewRouter()
	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "I am alive")
	})
}
