package main

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func color(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "BLUEE")
}

func main() {
	http.HandleFunc("/color", color)
	http.ListenAndServe(":8090", nil)

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":3000", nil)
}
