package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/FinweVI/dedibox-exporter/collectors"
	"github.com/FinweVI/dedibox-exporter/online"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	var (
		c   *online.Client
		err error

		myCollectors  collectorSlice
		listenAddress string
		metricsPath   string
		logLevel      int
	)

	flag.Var(&myCollectors, "collector", fmt.Sprintf("List of Collectors to enable (%s) (default \"abuse\")", strings.Join(validCollectors, ", ")))
	flag.StringVar(&listenAddress, "listen-address", "127.0.0.1:9539", "Address to listen on for web interface and telemetry")
	flag.StringVar(&metricsPath, "metric-path", "/metrics", "Path under which to expose metrics")
	flag.IntVar(&logLevel, "log-level", 1, "Log level: 0=debug, 1=info, 2=warn, 3=error")
	flag.Parse()

	setLogLevel(logLevel)

	if token, ok := os.LookupEnv("ONLINE_API_TOKEN"); !ok {
		log.Fatal("Please provide your API Token as an env var 'ONLINE_API_TOKEN'")
	} else {
		c, err = online.NewClient(online.AuthToken(token))
		if err != nil {
			log.Error("Unable to create API client")
			log.Debug(err)
		}
	}

	if len(myCollectors) == 0 {
		myCollectors.SetDefaultCollector()
		log.WithField("collectors", myCollectors).
			Debug("No collector selected, using default configuration")
	}

	for _, cltr := range myCollectors {
		switch cltr {
		case "abuse":
			prometheus.MustRegister(collectors.NewAbuseCollector(c))
		case "dedibackup":
			prometheus.MustRegister(collectors.NewDedibackupCollector(c))
		case "plan":
			prometheus.MustRegister(collectors.NewPlanCollector(c))
		case "ddos":
			prometheus.MustRegister(collectors.NewDDoSCollector(c))
		}
	}

	log.WithFields(log.Fields{
		"collectors": myCollectors,
		"address":    listenAddress,
	}).Info("Starting Dedibox Exporter")
	http.Handle(metricsPath, promhttp.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write(
			[]byte(
				`<html>
			     		<head><title>Dedibox Exporter</title></head>
			     		<body>
			     			<h1>Dedibox Exporter</h1>
			     			<p><a href='` + metricsPath + `'>Metrics</a></p>
			     		</body>
			     	</html>`,
			))
		if err != nil {
			log.Debug(err)
			log.WithField("path", "/").Error("unable to show the page")
		}
	})
	log.Fatal(http.ListenAndServe(listenAddress, nil))
}
