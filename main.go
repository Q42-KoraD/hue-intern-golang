package main

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func color(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "The current version of this server is: Blue \n")
	fmt.Fprintf(w, "For Prometheus metrics go to /metrics")
}

func main() {
	http.HandleFunc("/", color)
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":3000", nil)
}
