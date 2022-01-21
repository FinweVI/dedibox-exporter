package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/FinweVI/dedibox-exporter/collectors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	myCollectors  collectorSlice
	listenAddress string
	metricsPath   string
)

func main() {
	if _, ok := os.LookupEnv("ONLINE_API_TOKEN"); !ok {
		log.Fatalf("Please provide your API Token as an env var 'ONLINE_API_TOKEN'")
	}

	flag.Var(&myCollectors, "collector", fmt.Sprintf("List of Collectors to enable (%s) (default \"abuse\")", strings.Join(validCollectors, ", ")))
	flag.StringVar(&listenAddress, "listen-address", "127.0.0.1:9539", "Address to listen on for web interface and telemetry")
	flag.StringVar(&metricsPath, "metric-path", "/metrics", "Path under which to expose metrics")
	flag.Parse()

	if len(myCollectors) == 0 {
		myCollectors = append(myCollectors, "abuse")
	}

	for _, cltr := range myCollectors {
		switch cltr {
		case "abuse":
			prometheus.MustRegister(collectors.NewAbuseCollector())
		case "dedibackup":
			prometheus.MustRegister(collectors.NewServerCollector())
		case "plan":
			prometheus.MustRegister(collectors.NewPlanCollector())
		case "ddos":
			prometheus.MustRegister(collectors.NewDDoSCollector())
		}
	}

	log.Printf("Dedibox Exporter")
	log.Printf("Starting Server: %s", listenAddress)
	http.Handle(metricsPath, promhttp.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
             <head><title>Dedibox Exporter</title></head>
             <body>
             <h1>Dedibox Exporter</h1>
             <p><a href='` + metricsPath + `'>Metrics</a></p>
             </body>
             </html>`))
	})
	log.Fatal(http.ListenAndServe(listenAddress, nil))
}
