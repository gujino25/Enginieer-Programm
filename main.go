package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Tupe", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	})

	log.Println("Сервер запущен на: 9091")
	if err := http.ListenAndServe(":9091", mux); err != nil {
		log.Fatal(err)
	}
}
