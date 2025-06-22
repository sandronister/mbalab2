package main

import (
	"context"
	"errors"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/sandronister/mba-lab2/cep-service/internal/infra/web"
	"github.com/sandronister/mba-lab2/cep-service/internal/usecase"
	"github.com/sandronister/mba-lab2/pkg/otel"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() (err error) {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	otelShutdown, err := otel.SetupOTelSDK("cep-service", ctx)
	if err != nil {
		return
	}
	defer func() {
		err = errors.Join(err, otelShutdown(context.Background()))
	}()

	serverPort := os.Getenv("CEP_SERVICE_SERVER_PORT")
	if serverPort == "" {
		serverPort = ":8080"
	}
	srv := &http.Server{
		Addr:         serverPort,
		BaseContext:  func(_ net.Listener) context.Context { return ctx },
		ReadTimeout:  time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      newHTTPHandler(),
	}
	srvErr := make(chan error, 1)
	go func() {
		log.Printf("Starting server on port %s...", serverPort[1:])
		srvErr <- srv.ListenAndServe()
	}()

	select {
	case err = <-srvErr:
		return
	case <-ctx.Done():
		stop()
	}

	err = srv.Shutdown(context.Background())
	return
}

func newHTTPHandler() http.Handler {
	mux := http.NewServeMux()

	localeFinder := usecase.NewLocaleFinder(http.DefaultClient)
	mux.HandleFunc("POST /", web.NewLocaleHandler(localeFinder).Handle)

	return otelhttp.NewHandler(mux, "/")
}
