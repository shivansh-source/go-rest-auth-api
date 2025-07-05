package controllers

import (
	"net/http"
)

func ProtectedEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("🎉 Access granted to protected route!"))
}
