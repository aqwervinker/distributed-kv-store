package api

import (
	"encoding/json"
	"net/http"

	"distributed-kv-store/internal/consensus"
	"distributed-kv-store/internal/security"

	"github.com/gorilla/mux"
)

func SetupRouter(zab *consensus.Zab) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/login", security.Login).Methods("POST")

	s := r.PathPrefix("/api").Subrouter()
	s.Use(security.Authenticate)

	s.HandleFunc("/key", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "PUT":
			var kv map[string]string
			if err := json.NewDecoder(r.Body).Decode(&kv); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			for k, v := range kv {
				zab.Put(k, v)
			}
			w.WriteHeader(http.StatusNoContent)
		case "GET":
			key := r.URL.Query().Get("key")
			value, ok := zab.Get(key)
			if !ok {
				http.Error(w, "key not found", http.StatusNotFound)
				return
			}
			json.NewEncoder(w).Encode(map[string]string{key: value})
		case "DELETE":
			key := r.URL.Query().Get("key")
			zab.Delete(key)
			w.WriteHeader(http.StatusNoContent)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	}).Methods("PUT", "GET", "DELETE")

	return r
}
