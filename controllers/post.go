package controllers

import (
	"encoding/json"
	"net/http"
     "github.com/gorilla/mux"
	"github.com/shivansh-source/go-rest-auth-api/config"
	"github.com/shivansh-source/go-rest-auth-api/models"
)

// CreatePost - POST /api/posts
func CreatePost(w http.ResponseWriter, r *http.Request) {
	var post models.Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// For now, assume UserID is provided in JSON; later get from JWT claims
	if err := config.DB.Create(&post).Error; err != nil {
		http.Error(w, "Failed to create post", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
jsonData, _ := json.MarshalIndent(post, "", "  ")
w.Write(jsonData)

}

// GetAllPosts - GET /api/posts
func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	var posts []models.Post
	config.DB.Find(&posts)
	json.NewEncoder(w).Encode(posts)
}

// UpdatePost - PUT /api/posts/{id}
func UpdatePost(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var post models.Post

	if err := config.DB.First(&post, id).Error; err != nil {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}

	var updated models.Post
	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	post.Title = updated.Title
	post.Content = updated.Content

	config.DB.Save(&post)
	w.Header().Set("Content-Type", "application/json")
jsonData, _ := json.MarshalIndent(post, "", "  ")
w.Write(jsonData)

}

// DeletePost - DELETE /api/posts/{id}
func DeletePost(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if err := config.DB.Delete(&models.Post{}, id).Error; err != nil {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
// GetPostByID - GET /api/posts/{id}
func GetPostByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var post models.Post

	if err := config.DB.First(&post, id).Error; err != nil {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
jsonData, _ := json.MarshalIndent(post, "", "  ")
w.Write(jsonData)

}			