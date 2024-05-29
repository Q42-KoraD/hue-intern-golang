package main

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func color(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "BLUEEeee")
}

func main() {
	http.HandleFunc("/color", color)
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":3000", nil)
}
