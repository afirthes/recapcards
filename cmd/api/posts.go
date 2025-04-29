package main

import (
	"errors"
	"github.com/afirthes/recapcards/internal/store"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"strconv"
)

type CreatePostPayload struct {
	Title   string   `json:"title" validate:"required,max=100"`
	Content string   `json:"content" validate:"required,max=1000"`
	Tags    []string `json:"tags"`
}

func (app *application) createPostHandler(w http.ResponseWriter, r *http.Request) {

	var postDTO CreatePostPayload
	if err := readJSON(w, r, &postDTO); err != nil {
		writeJSONError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	post := &store.Post{
		Title:   postDTO.Title,
		Content: postDTO.Content,
		UserID:  1,
		Tags:    postDTO.Tags,
	}

	if err := app.storage.Posts.Create(r.Context(), post); err != nil {
		log.Println(err.Error())
		writeJSONError(w, http.StatusInternalServerError, "Error creating post id db")
		return
	}

	if err := writeJSON(w, http.StatusCreated, post); err != nil {
		log.Println(err.Error())
		writeJSONError(w, http.StatusInternalServerError, "Error writing post to response")
		return
	}

}

func (app *application) getPostHandler(w http.ResponseWriter, r *http.Request) {
	postID := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(postID, 10, 64)
	if err != nil {
		log.Println(err.Error())
		writeJSONError(w, http.StatusBadRequest, "Invalid post id")
		return
	}

	post, err := app.storage.Posts.GetByID(r.Context(), id)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, store.ErrNotFound) {
			writeJSONError(w, http.StatusNotFound, "Post not found")
			return
		} else {
			writeJSONError(w, http.StatusInternalServerError, "Error getting post")
			return
		}
	}

	if err := writeJSON(w, http.StatusOK, post); err != nil {
		log.Println(err.Error())
		writeJSONError(w, http.StatusInternalServerError, "Error writing post to response")
		return
	}
}
