package main

import (
	"encoding/json"
	"log"
	"net/http"

	expand "github.com/openvenues/gopostal/expand"
	parser "github.com/openvenues/gopostal/parser"
)

type Input struct {
	Address string `json:address`
}

func jsonErrorMessage(msg string) []byte {
	r := make(map[string]string)
	r["error"] = msg
	j, _ := json.Marshal(r)
	return j
}

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

func expandHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		w.Header().Set("Content-Type", "application/json")

		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(w, string(jsonErrorMessage("Content-Type must be application/json")), 400)
			return
		}

		var input Input

		err := json.NewDecoder(r.Body).Decode(&input)

		if err != nil {
			http.Error(w, string(jsonErrorMessage("error decoding request")), 500)
			return
		}

		w.WriteHeader(http.StatusOK)

		expansions := expand.ExpandAddress(input.Address)
		expansionsJson, _ := json.Marshal(expansions)

		w.Write(expansionsJson)

		return
	default:
		http.Error(w, "405 method not allowed", 405)
		return
	}
}

func parseHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		w.Header().Set("Content-Type", "application/json")

		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(w, string(jsonErrorMessage("Content-Type must be application/json")), 400)
			return
		}

		var input Input

		err := json.NewDecoder(r.Body).Decode(&input)

		if err != nil {
			http.Error(w, string(jsonErrorMessage("error decoding request")), 500)
			return
		}

		w.WriteHeader(http.StatusOK)

		parsed := make(map[string]string)
		components := parser.ParseAddress(input.Address)

		for _, component := range components {
			parsed[component.Label] = component.Value
		}

		parsedJson, _ := json.Marshal(parsed)

		w.Write(parsedJson)

		return
	default:
		http.Error(w, "405 method not allowed", 405)
		return
	}
}

func main() {
	http.HandleFunc("/api/v1/expand", expandHandler)
	http.HandleFunc("/api/v1/parse", parseHandler)
	http.HandleFunc("/healthz", healthHandler)

	log.Fatal(http.ListenAndServe(":8123", nil))
}
