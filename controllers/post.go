package controllers

import (
    "net/http"
)

// Example placeholder function
func GetAllPosts(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("This will return all posts"))
}
