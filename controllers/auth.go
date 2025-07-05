package controllers

import (
    "net/http"
    "encoding/json"

    "github.com/shivansh-source/go-rest-auth-api/config"
    "github.com/shivansh-source/go-rest-auth-api/models"
    "github.com/shivansh-source/go-rest-auth-api/utils"
)

func Signup(w http.ResponseWriter, r *http.Request) {
    var input struct {
        Name     string `json:"name"`
        Email    string `json:"email"`
        Password string `json:"password"`
    }

    // Parse request body
    if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    // Hash the password
    hashedPassword, err := utils.HashPassword(input.Password)
    if err != nil {
        http.Error(w, "Failed to hash password", http.StatusInternalServerError)
        return
    }

    // Create user record
    user := models.User{
        Name:     input.Name,
        Email:    input.Email,
        Password: hashedPassword,
    }

    if err := config.DB.Create(&user).Error; err != nil {
        http.Error(w, "Email already taken or failed to create user", http.StatusBadRequest)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]string{"message": "Signup successful"})
}
func Login(w http.ResponseWriter, r *http.Request) {
    var input struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }

    if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    var user models.User
    result := config.DB.First(&user, "email = ?", input.Email)
    if result.Error != nil {
        http.Error(w, "User not found", http.StatusUnauthorized)
        return
    }

    // Compare the password
    if !utils.CheckPasswordHash(input.Password, user.Password) {
        http.Error(w, "Invalid password", http.StatusUnauthorized)
        return
    }

    // Generate JWT token
    token, err := utils.GenerateJWT(user.ID)
    if err != nil {
        http.Error(w, "Failed to generate token", http.StatusInternalServerError)
        return
    }

    // Return token
    json.NewEncoder(w).Encode(map[string]string{
        "token": token,
    })
}

