package main

import (
	"net/http"

	"github.com/pbaettig/promscriptsgw/internal/app"
	"github.com/pbaettig/promscriptsgw/internal/cfg"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	log "github.com/sirupsen/logrus"
)

func main() {
	err := cfg.FromCommandline()
	if err != nil {
		log.Fatal(err.Error())
	}

	if cfg.Debug {
		log.SetLevel(log.DebugLevel)
	}

	log.Infof("looking for executable scripts in \"%s\"", cfg.ScriptDir)

	app.CollectionLoop()
	http.Handle(cfg.MetricsURL, promhttp.HandlerFor(app.Registry, promhttp.HandlerOpts{
		ErrorLog:      log.StandardLogger(),
		ErrorHandling: 1,
	}))

	log.Infof("listening for requests on %s", cfg.ListenAddr)
	log.Fatal(http.ListenAndServe(cfg.ListenAddr, nil))
}
