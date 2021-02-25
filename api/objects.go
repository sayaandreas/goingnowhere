package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/sayaandreas/goingnowhere/storage"
)

var objectKey = "objectKey"

func objects(router chi.Router) {
	router.Post("/", upload)
	router.Route("/{objectKey}", func(router chi.Router) {
		router.Use(ObjectContext)
		router.Get("/", getAllObjects)
	})
}

func ObjectContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		objectKeyParam := chi.URLParam(r, "objectKey")
		if objectKeyParam == "" {
			render.Render(w, r, ErrorRenderer(fmt.Errorf("object key is required")))
			return
		}
		ctx := context.WithValue(r.Context(), objectKey, objectKeyParam)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func upload(w http.ResponseWriter, r *http.Request) {
	var ur storage.UploadRequest
	err := json.NewDecoder(r.Body).Decode(&ur)
	if err != nil {
		fmt.Fprint(w, err)
	}
	resp, err := storageInstance.UploadFile(ur)
	if err != nil {
		fmt.Fprint(w, err)
	}
	fmt.Fprint(w, resp)
}
