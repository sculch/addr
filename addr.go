package main

import (
	"log"
	"net/http"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		return
	default:
		http.Error(w, "405 method not allowed", 405)
		return
	}
}

func main() {
	// http.HandleFunc("/api/v1/expand", expandHandler)
	// http.HandleFunc("/api/v1/parse", parseHandler)
	http.HandleFunc("/healthz", healthHandler)

	log.Fatal(http.ListenAndServe(":8123", nil))
}
