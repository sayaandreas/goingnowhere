package api

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

var s3IDKey = "s3ID"

func s3(router chi.Router) {
	router.Get("/", getAllPlaces)
}

func getAllPlaces(w http.ResponseWriter, r *http.Request) {
	resp := storageInstance.GetBucketList()
	fmt.Fprint(w, resp)
}
