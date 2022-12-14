package handlers

// import (
// 	"encoding/json"
// 	"net/http"
// 	"time"

// 	"github.com/batistondeoliveira/fullcycle_curso_golang/09-APIS/internal/dto"
// 	"github.com/batistondeoliveira/fullcycle_curso_golang/09-APIS/internal/entity"
// 	"github.com/batistondeoliveira/fullcycle_curso_golang/09-APIS/internal/infra/database"
// 	"github.com/go-chi/jwtauth"
// )

// type UserHandler struct {
// 	UserDB       database.UserInterface
// 	Jwt          *jwtauth.JWTAuth
// 	JwtExpiresIn int
// }

// func NewUserHandler(userDB database.UserInterface, jwt *jwtauth.JWTAuth, jwtExpiresIn int) *UserHandler {
// 	return &UserHandler{
// 		UserDB:       userDB,
// 		Jwt:          jwt,
// 		JwtExpiresIn: jwtExpiresIn,
// 	}
// }

// func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
// 	var user dto.CreateUserInput
// 	err := json.NewDecoder(r.Body).Decode(&user)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}
// 	u, err := entity.NewUser(user.Name, user.Email, user.Password)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}
// 	err = h.UserDB.Create(u)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
// 	w.WriteHeader(http.StatusCreated)
// }

// func (h *UserHandler) GetJWT(w http.ResponseWriter, r *http.Request) {
// 	var user dto.GetJWTInput
// 	err := json.NewDecoder(r.Body).Decode(&user)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}
// 	u, err := h.UserDB.FindByEmail(user.Email)
// 	if err != nil {
// 		w.WriteHeader(http.StatusUnauthorized)
// 		return
// 	}
// 	if !u.ValidatePassword(user.Password) {
// 		w.WriteHeader(http.StatusUnauthorized)
// 	}
// 	_, tokenString, _ := h.Jwt.Encode(map[string]interface{}{
// 		"sub": u.ID.String(),
// 		"exp": time.Now().Add(time.Second * time.Duration(h.JwtExpiresIn)).Unix(),
// 	})
// 	accessToken := struct {
// 		AccessToken string `json:"access_token"`
// 	}{
// 		AccessToken: tokenString,
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(accessToken)
// }
