package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	cfg := readCfg()

	proxy := newProxy(cfg)

	http.HandleFunc("/api/movies", proxy.proxy)
	http.HandleFunc("/health", handleHealth)

	port := cfg.PORT
	if port == "" {
		port = "8081"
	}

	log.Printf("cfg %+v", cfg)

	log.Printf("Starting proxy microservice on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]bool{"status": true})
}
