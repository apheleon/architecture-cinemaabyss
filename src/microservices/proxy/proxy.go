package main

import (
	"io"
	"log"
	"math/rand"
	"net/http"
	"net/url"
)

type proxy interface {
	proxy(w http.ResponseWriter, r *http.Request)
}

type BaseProxy struct {
	URL string
}

type StranglerFigProxy struct {
	Enabled          bool
	MigrationPercent int
	MonolithProxy    proxy
	ServiceProxy     proxy
}

func (b BaseProxy) proxy(w http.ResponseWriter, r *http.Request) {
	proxyUrl, _ := url.Parse(b.URL)

	proxyUrl.Path = r.URL.Path
	proxyUrl.RawQuery = r.URL.RawQuery

	req, err := http.NewRequest(r.Method, proxyUrl.String(), r.Body)
	if err != nil {
		log.Printf("Error during NewRequest() %s: %s\n", proxyUrl.String(), err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	for key, values := range r.Header {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("Error during Do() %s: %s\n", proxyUrl.String(), err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	written, err := io.Copy(w, resp.Body)
	if err != nil {
		log.Printf("Error during Copy() %s: %s\n", proxyUrl.String(), err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("%s - %s - %s - %d - %dKB\n", r.Proto, r.Method, proxyUrl.String(), resp.StatusCode, written/1000)
}

func (b StranglerFigProxy) proxy(w http.ResponseWriter, r *http.Request) {
	decorate := b.MonolithProxy

	if b.Enabled && (b.MigrationPercent > 0 && rand.Intn(100) <= b.MigrationPercent) {
		decorate = b.ServiceProxy
	}

	decorate.proxy(w, r)
}

func newBaseProxy(url string) BaseProxy {
	return BaseProxy{URL: url}
}

func newStranglerFigProxy(enabled bool, migrationPercent int, monolithProxy, serviceProxy proxy) StranglerFigProxy {
	return StranglerFigProxy{Enabled: enabled, MigrationPercent: migrationPercent, MonolithProxy: monolithProxy, ServiceProxy: serviceProxy}
}

func newProxy(cfg config) proxy {
	return newStranglerFigProxy(cfg.GRADUAL_MIGRATION, cfg.MOVIES_MIGRATION_PERCENT, newBaseProxy(cfg.MONOLITH_URL), newBaseProxy(cfg.MOVIES_SERVICE_URL))
}
