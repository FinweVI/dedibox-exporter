package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/FinweVI/dedibox-exporter/collectors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	listenAddress := flag.String("listen-address", ":9539", "Address to listen on for web interface and telemetry.")
	metricsPath := flag.String("metric-path", "/metrics", "Path under which to expose metrics.")
	flag.Parse()

	if _, ok := os.LookupEnv("ONLINE_API_TOKEN"); !ok {
		log.Fatalf("Please provide your API Token as an env var 'ONLINE_API_TOKEN'")
	}

	prometheus.MustRegister(collectors.NewAbuseCollector())
	prometheus.MustRegister(collectors.NewServerCollector())
	prometheus.MustRegister(collectors.NewPlanCollector())

	log.Printf("Dedibox Exporter")
	log.Printf("Starting Server: %s", *listenAddress)
	http.Handle(*metricsPath, promhttp.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
             <head><title>Dedibox Exporter</title></head>
             <body>
             <h1>Dedibox Exporter</h1>
             <p><a href='` + *metricsPath + `'>Metrics</a></p>
             </body>
             </html>`))
	})
	log.Fatal(http.ListenAndServe(*listenAddress, nil))
}
