package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func main() {
	http.Handle("/metrics", promhttp.Handler()) // endpoint to expose prometheus formatted metrics
	http.HandleFunc("/", homeHandler)
	log.Fatal(http.ListenAndServe(":8082", nil))
}
