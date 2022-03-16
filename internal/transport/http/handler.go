package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/PatriciaChebet/rest_api_with_go/internal/comment"
	"github.com/gorilla/mux"
)

//Handler - stores pointer to our comment service
type Handler struct {
	Router  *mux.Router
	Service *comment.Service
}

type Response struct {
	Message string
}

//NewHandler - returns pointer to a Handler
func NewHandler(service *comment.Service) *Handler {
	return &Handler{
		Service: service,
	}
}

//SetUpRoutes - sets up all the routes for our application
func (h *Handler) SetUpRoutes() {
	fmt.Println("Setting up Routes")
	h.Router = mux.NewRouter()
	h.Router.HandleFunc("/api/comment", h.GetAllComments).Methods("GET")
	h.Router.HandleFunc("/api/comment/{id}", h.GetComment).Methods("GET")
	h.Router.HandleFunc("/api/comment", h.PostComment).Methods("POST")
	h.Router.HandleFunc("/api/comment/{id}", h.DeleteComment).Methods("DELETE")
	h.Router.HandleFunc("/api/comment/{id}", h.UpdateComment).Methods("PUT")
	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json, charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(Response{Message: "I am alive"}); err != nil {
			panic(err)
		}
	})
}

//GetComment - retrieve comment by id
func (h *Handler) GetComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json, charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	vars := mux.Vars(r)
	id := vars["id"]

	i, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		fmt.Fprintf(w, "Unable to parse UINT from ID")
	}

	comment, err := h.Service.GetComment(uint(i))
	if err != nil {
		fmt.Fprintf(w, "Error retrieving Comment by ID")
	}
	if err := json.NewEncoder(w).Encode(comment); err != nil {
		panic(err)
	}
}

//GetAllComments - retrieves all the comments from the comment service
func (h *Handler) GetAllComments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json, charset = UTF-8")
	w.WriteHeader(http.StatusOK)
	comments, err := h.Service.GetAllComments()
	if err != nil {
		fmt.Fprintf(w, "Failed to retrieve all comments")
	}
	if err := json.NewEncoder(w).Encode(comments); err != nil {
		panic(err)
	}
}

//PostComment - adds a new comment
func (h *Handler) PostComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json, charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	var comment comment.Comment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		fmt.Fprintf(w, "failed to decode JSON body")
	}
	comment, err := h.Service.PostComment(comment)
	if err != nil {
		fmt.Fprintf(w, "Failed to post new comment")
	}
	if err := json.NewEncoder(w).Encode(comment); err != nil {
		panic(err)
	}
}

//UpdateComment - updates comment by id
func (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json, charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	var comment comment.Comment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		fmt.Fprintf(w, "failed to decode JSON body")
	}

	vars := mux.Vars(r)
	id := vars["id"]

	commentID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		fmt.Fprintf(w, "Failed to parse uint from ID")
	}

	comment, err = h.Service.UpdateComment(uint(commentID), comment)
	if err != nil {
		fmt.Fprintf(w, "Failed to update comment")
	}
	if err := json.NewEncoder(w).Encode(comment); err != nil {
		panic(err)
	}
}

//DeleteComment - deletes a comment by id
func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json, charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	id := vars["id"]

	commentID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		fmt.Fprintf(w, "Failed to parse uint from ID")
	}

	err = h.Service.DeleteComment(uint(commentID))
	if err != nil {
		fmt.Fprintf(w, "failed to delete comment by comment ID")
	}

	if err := json.NewEncoder(w).Encode(Response{Message: "Comment successfully deleted"}); err != nil {
		panic(err)
	}
}
