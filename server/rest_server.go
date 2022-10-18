package server

import (
	"fmt"

	"net/http"
)

func ServeProduct(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		get(w, r)
	case http.MethodPost:
		post(w, r)
	case http.MethodDelete:
		delete(w, r)
	case http.MethodPut:
		put(w, r)
	}
}

func get(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HelloFromGet")
}

func put(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HelloFromPutID")
}

func post(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HelloFromRaz")
}

func delete(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HelloFromDeletedID")
}

func SetupServer() {
	port := ":8080"

	http.ListenAndServe(port, http.HandlerFunc(ServeProduct))
}
