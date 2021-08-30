package main

import (
	"log"
	"net/http"
)

func logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s", r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
