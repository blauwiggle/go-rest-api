package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/blauwiggle/go-rest-api/internal/comment"
	"github.com/gorilla/mux"
)

// Handler - stores pointer to our comments service
type Handler struct {
	Router  *mux.Router
	Service *comment.Service
}

// Respone - returns a response
type Respone struct {
	Message string
}

// NewHandler - returns a pointer to a new Handler
func NewHandler(service *comment.Service) *Handler {
	return &Handler{
		Service: service,
	}
}

// SetupRoute - sets up our routes
func (h *Handler) SetupRoutes() {
	fmt.Println("Setting up routes")

	h.Router = mux.NewRouter()

	h.Router.HandleFunc("/api/comment", h.GetAllComments).Methods("GET")
	h.Router.HandleFunc("/api/comment", h.PostComment).Methods("POST")
	h.Router.HandleFunc("/api/comment/{id}", h.GetComment).Methods("GET")
	h.Router.HandleFunc("/api/comment/{id}", h.UpdateComment).Methods("PUT")
	h.Router.HandleFunc("/api/comment/{id}", h.DeleteComment).Methods("DELETE")

	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(Respone{Message: "I am Alive"}); err != nil {
			panic(err)
		}
	})

}

// GetComment - gets a comment by id
func (h *Handler) GetComment(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	id := vars["id"]

	i, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		fmt.Println("Failed to parse id")
		fmt.Println(err)
	}

	comment, err := h.Service.GetComment(uint(i))

	if err != nil {
		fmt.Println("Error getting comment")
		fmt.Println(err)
	}

	if err := json.NewEncoder(w).Encode(comment); err != nil {
		panic(err)
	}
}

// GetAllComments - gets all comments
func (h *Handler) GetAllComments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	comments, err := h.Service.GetAllComments()

	if err != nil {
		fmt.Println(w, "Failed to retrieve all comments")
	}

	if err := json.NewEncoder(w).Encode(comments); err != nil {
		panic(err)
	}

}

// PostComment - posts a comment
func (h *Handler) PostComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	var comment comment.Comment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		fmt.Fprintf(w, "Error decoding comment")
	}

	comment, err := h.Service.PostComment(comment)
	if err != nil {
		fmt.Println(w, "Failed to posting comment")
	}

	if err := json.NewEncoder(w).Encode(comment); err != nil {
		panic(err)
	}
}

// UpdateComment - updates a comment
func (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	id := vars["id"]

	i, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		fmt.Println(w, "Failed to parse id")
	}

	var comment comment.Comment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		fmt.Fprintf(w, "Error decoding comment")
	}

	updatedComment, err := h.Service.UpdateComment(uint(i), comment)

	if err := json.NewEncoder(w).Encode(updatedComment); err != nil {
		panic(err)
	}

}

// DeleteComment - deletes a comment
func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	vars := mux.Vars(r)
	id := vars["id"]

	i, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		fmt.Println(w, "Failed to parse id")
	}

	err = h.Service.DeleteComment(uint(i))

	if err != nil {
		fmt.Println("Error deleting comment")
		fmt.Println(err)
	}

	if err := json.NewEncoder(w).Encode("Comment successfully deleted"); err != nil {
		panic(err)
	}
}
