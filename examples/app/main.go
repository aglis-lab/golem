package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/aglis-lab/golem/src/app"
	"github.com/aglis-lab/golem/src/middleware"
	v1 "github.com/aglis-lab/golem/src/v1"

	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
)

func main() {
	// Handle SIGINT (CTRL+C) gracefully.
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// Init app context
	if err := app.Init(ctx); err != nil {
		log.Panic(err)
	}

	// TODO: Add Otel SDK
	// // Setup Otel SDK
	// otelShutdown, err := tracer.SetupOTelSDK(ctx, app.Config().OltpGRPCProvider, app.Config().ServiceName, app.Config().ServiceVersion)
	// if err != nil {
	// 	log.Panic(err)
	// }

	// // Handle shutdown properly so nothing leaks.
	// defer func() {
	// 	err = errors.Join(err, otelShutdown(context.Background()))
	// 	log.Println(err)
	// }()

	// Init Router
	address := fmt.Sprintf(":%d", app.Config().BindAddress)
	router := initRouter(ctx, address)

	// TODO: Active Tracer
	// Start HTTP Server
	srv := &http.Server{
		Addr:         address,
		BaseContext:  func(_ net.Listener) context.Context { return ctx },
		ReadTimeout:  time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      router,
	}
	srvErr := make(chan error, 1)

	log.Printf("Starting service on %s", address)
	go func() {
		srvErr <- srv.ListenAndServe()
	}()

	// Wait for interruption.
	select {
	case err := <-srvErr:
		// Error when starting HTTP server.
		fmt.Fprintf(os.Stderr, "Error starting HTTP server: %v\n", err)
		return
	case <-ctx.Done():
		// Wait for first CTRL+C.
		// Stop receiving signal notifications as soon as possible.
		// stop()
	}

	// When Shutdown is called, ListenAndServe immediately returns ErrServerClosed.
	err := srv.Shutdown(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error shutting down HTTP server: %v\n", err)
	}
}

func initRouter(ctx context.Context, address string) *chi.Mux {
	r := chi.NewRouter()
	r.Use(chimiddleware.Recoverer)
	r.Use(middleware.RequestIDContext(middleware.DefaultGenerator))
	r.Use(middleware.RequestAttributesContext)
	r.Use(chimiddleware.Logger)
	r.Use(chimiddleware.RealIP)
	r.Use(chimiddleware.Timeout(60 * time.Second))

	deps := v1.Dependencies(ctx)

	r.Route("/v1", func(r chi.Router) {
		v1.Router(r, deps)
	})

	return r
}
