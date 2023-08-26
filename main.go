package main

import (
	"fmt"
	"net/http"
	"os"

	"stack-test-todolist/router"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Logger struct{ handler http.Handler }

func (l Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Info().
		Str("method", r.Method).
		Str("uri", r.URL.String()).
		Str("ipaddr", r.RemoteAddr).
		Str("referer", r.Header.Get("Referer")).
		Str("userAgent", r.Header.Get("User-Agent")).
		Msg("")
	l.handler.ServeHTTP(w, r)
}

func main() {
	router := router.New()
	host := "localhost"
	port := 3000
	addr := fmt.Sprintf("%v:%d", host, port)
	handler := Logger{router}

	// zerolog setup
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	log.Info().Str("addr", addr).Msg("🚀 Started Server")
	log.Error().Err(http.ListenAndServe(addr, handler)).Msg("")
}
