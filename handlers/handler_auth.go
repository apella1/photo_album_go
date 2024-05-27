package handlers

import (
	"album/internal/database"
	"album/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No environment variables file found")
	}
}

func ComparePasswords(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func (h *Handler) LoginUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Email    string
		Password string
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}
	user, err := h.Cfg.DB.GetUserByEmail(r.Context(), params.Email)
	if err != nil {
		utils.RespondWithError(w, 400, "Invalid credentials!")
		return
	}
	if !ComparePasswords(user.Password, params.Password) {
		utils.RespondWithError(w, 400, "Invalid credentials")
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
	})
	hmacSampleSecret := []byte(os.Getenv("HMAC_SECRET_KEY"))
	tokenStr, err := token.SignedString(hmacSampleSecret)
	if err != nil {
		utils.RespondWithError(w, 500, "Couldn't generate token")
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"token": tokenStr,
	})
}

func (h *Handler) GetUserByJWT(w http.ResponseWriter, r *http.Request, user database.User) {
	utils.RespondWithJSON(w, 200, utils.DatabaseUserToUser(user))
}
