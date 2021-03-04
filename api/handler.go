package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/sayaandreas/goingnowhere/db"
	"github.com/sayaandreas/goingnowhere/storage"
)

var storageInstance storage.Storage
var enforcer *casbin.Enforcer
var dbInstance db.Database

func NewHandler(s storage.Storage, e *casbin.Enforcer, d db.Database) http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	storageInstance = s
	enforcer = e
	dbInstance = d
	router.MethodNotAllowed(methodNotAllowedHandler)
	router.NotFound(notFoundHandler)
	router.Route("/buckets", bucket)
	router.Route("/objects", objects)
	router.Route("/auth", auth)
	router.Post("/resources", resources)
	return router
}
func methodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(405)
	render.Render(w, r, ErrMethodNotAllowed)
}
func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(400)
	render.Render(w, r, ErrNotFound)
}

type ResourceRequest struct {
	Sub string `json:"sub"`
	Obj string `json:"obj"`
	Act string `json:"act"`
}

func resources(w http.ResponseWriter, r *http.Request) {
	var ro ResourceRequest
	err := json.NewDecoder(r.Body).Decode(&ro)
	if err != nil {
		fmt.Println(err)
	}
	ok, err := enforcer.Enforce(ro.Sub, ro.Obj, ro.Act)
	if err != nil {
		render.Render(w, r, ErrForbiden)
		return
	}
	if !ok {
		render.Render(w, r, ErrUnauthorized)
		return
	}
	fmt.Fprintln(w, "Authz")
}
