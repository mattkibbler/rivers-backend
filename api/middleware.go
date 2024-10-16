package api

import (
	"fmt"
	"net/http"
)

func ApiGet(server *ApiServer, handler http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "status not allowed")
			return
		}
		handler(w, r)
	})
}

func ApiPost(server *ApiServer, handler http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "status not allowed")
			return
		}
		handler(w, r)
	})
}
