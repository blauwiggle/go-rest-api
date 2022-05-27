package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/blauwiggle/go-rest-api/internal/comment"
	"github.com/gorilla/mux"
)

// GetComment - gets a comment by id
func (h *Handler) GetComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	i, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		sendErrorResponse(w, "Failed to parse id", err)
		return
	}

	comment, err := h.Service.GetComment(uint(i))

	if err != nil {
		sendErrorResponse(w, "Error getting comment", err)
		return
	}

	if err := sendOkResponse(w, comment); err != nil {
		panic(err)
	}
}

// GetAllComments - gets all comments
func (h *Handler) GetAllComments(w http.ResponseWriter, r *http.Request) {
	comments, err := h.Service.GetAllComments()

	if err != nil {
		sendErrorResponse(w, "Failed to retrieve all comments", err)
		return
	}

	if err := sendOkResponse(w, comments); err != nil {
		panic(err)
	}

}

// PostComment - posts a comment
func (h *Handler) PostComment(w http.ResponseWriter, r *http.Request) {
	var comment comment.Comment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		sendErrorResponse(w, "Error decoding comment", err)
	}

	comment, err := h.Service.PostComment(comment)
	if err != nil {
		sendErrorResponse(w, "Failed to posting comment", err)
		return
	}

	if err := sendOkResponse(w, comment); err != nil {
		panic(err)
	}
}

// UpdateComment - updates a comment
func (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	i, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		sendErrorResponse(w, "Failed to parse id", err)
		return
	}

	var comment comment.Comment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		sendErrorResponse(w, "Error decoding comment", err)
		return
	}

	updatedComment, err := h.Service.UpdateComment(uint(i), comment)

	if err := sendOkResponse(w, updatedComment); err != nil {
		panic(err)
	}

}

// DeleteComment - deletes a comment
func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	i, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		sendErrorResponse(w, "Failed to parse id", err)
		return
	}

	err = h.Service.DeleteComment(uint(i))

	if err != nil {
		sendErrorResponse(w, "Error deleting comment", err)
		return
	}

	if err = sendOkResponse(w, "Comment deleted"); err != nil {
		panic(err)
	}
}
