package http

import (
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
		fmt.Fprintf(w, "I am alive")
	})

}

// GetComment - gets a comment by id
func (h *Handler) GetComment(w http.ResponseWriter, r *http.Request) {
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

	fmt.Fprintf(w, "%+v", comment)
}

// GetAllComments - gets all comments
func (h *Handler) GetAllComments(w http.ResponseWriter, r *http.Request) {
	comments, err := h.Service.GetAllComments()

	if err != nil {
		fmt.Println(w, "Failed to retrieve all comments")
	}

	fmt.Fprintf(w, "%+v", comments)
}

// PostComment - posts a comment
func (h *Handler) PostComment(w http.ResponseWriter, r *http.Request) {
	newComment, err := h.Service.PostComment(comment.Comment{
		Slug: "/",
	})

	if err != nil {
		fmt.Println(w, "Failed to posting comment")
	}

	fmt.Fprintf(w, "%+v", newComment)
}

// UpdateComment - updates a comment
func (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	updatedComment, err := h.Service.UpdateComment(1, comment.Comment{
		Slug: "/new",
	})

	if err != nil {
		fmt.Println(w, "Failed updating comment")
	}

	fmt.Fprintf(w, "%+v", updatedComment)

}

// DeleteComment - deletes a comment
func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {
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

	fmt.Fprintf(w, "Deleted comment")
}
