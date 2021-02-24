package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

var s3IDKey = "bucketName"

func s3(router chi.Router) {
	router.Get("/", getAllBuckets)
	router.Route("/{bucketName}", func(router chi.Router) {
		router.Use(APIContext)
		router.Get("/", getAllObjects)
	})

}

func APIContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bucketName := chi.URLParam(r, "bucketName")
		if bucketName == "" {
			render.Render(w, r, ErrorRenderer(fmt.Errorf("user ID is required")))
			return
		}
		ctx := context.WithValue(r.Context(), s3IDKey, bucketName)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getAllBuckets(w http.ResponseWriter, r *http.Request) {
	resp := storageInstance.GetBucketList()
	fmt.Fprint(w, resp)
}

func getAllObjects(w http.ResponseWriter, r *http.Request) {
	// bucketName := chi.URLParam(r, "bucketName")
	bucketName := r.Context().Value(s3IDKey).(string)
	resp := storageInstance.GetBucketObjectList(bucketName)
	fmt.Fprint(w, resp)
}
