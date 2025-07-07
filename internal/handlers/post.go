package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/engWaliullah/blog-server/internal/db"
	"github.com/engWaliullah/blog-server/internal/models"
	"github.com/engWaliullah/blog-server/internal/utils"
)

func GetPosts(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query("SELECT id, title, content FROM posts")
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "DB query failed")
		return
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		if err := rows.Scan(&post.ID, &post.Title, &post.Content); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, "Failed to read row")
			return
		}
		posts = append(posts, post)
	}

	utils.RespondWithJSON(w, http.StatusOK, posts)
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	var post models.Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request")
		return
	}

	err := db.DB.QueryRow(
		"INSERT INTO posts (title, content) VALUES ($1, $2) RETURNING id",
		post.Title, post.Content,
	).Scan(&post.ID)

	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Insert failed")
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, post)
}
