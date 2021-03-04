package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/sayaandreas/goingnowhere/jwe"
)

type LoginPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LogoutPayload struct {
	Token string `json:"token"`
}

func auth(router chi.Router) {
	router.Post("/login", login)
	router.Post("/logout", logout)
}

func login(w http.ResponseWriter, r *http.Request) {
	var lp LoginPayload
	err := json.NewDecoder(r.Body).Decode(&lp)
	if err != nil {
		fmt.Fprint(w, err)
	}
	user, err := dbInstance.GetUserByUsername(lp.Username)
	fmt.Println(user)
	if err != nil {
		fmt.Fprint(w, err)
	}
	pri := jwe.PrivateClaims{
		UserID: user.ID,
	}
	token, err := jwe.Encode(pri)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	fmt.Fprint(w, token)
}

func logout(w http.ResponseWriter, r *http.Request) {
	var lp LogoutPayload
	err := json.NewDecoder(r.Body).Decode(&lp)
	if err != nil {
		fmt.Fprint(w, err)
	}

	pub, pri, err := jwe.Decode(lp.Token)
	if err != nil {
		fmt.Fprint(w, err)
	}
	fmt.Fprintf(w, "Expiry: %d, ID: %d", pub.Expiry, pri.UserID)
}
