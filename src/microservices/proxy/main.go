package main

import (
	"log"
	"net/http"
)

func main() {
	cfg := readCfg()

	proxy := newProxy(cfg)

	http.HandleFunc("/api/movies", proxy.proxy)

	port := cfg.PORT
	if port == "" {
		port = "8081"
	}
	log.Printf("Starting proxy microservice on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
