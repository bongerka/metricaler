package app

import (
	"context"
	"errors"
	"github.com/bongerka/metricaler/internal/api"
	"github.com/bongerka/metricaler/internal/api/implementation"
	metimpl "github.com/bongerka/metricaler/internal/api/implementation/metric"
	"github.com/bongerka/metricaler/internal/config"
	metrepo "github.com/bongerka/metricaler/internal/repository/hashmap/metric"
	"gitlab.com/bongerka/lg"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

type App struct {
	cfg *config.Config
	l   lg.Logger

	httpServer *http.Server
	impl       *implementation.Implemetation
}

func NewApp(cfg *config.Config) *App {
	app := &App{cfg: cfg}

	repo := metrepo.NewRepository(cfg.Repo.Size)
	ms := metimpl.NewService(repo)
	app.impl = implementation.NewImplementation(ms)

	app.httpServer = app.initHTTPServer()
	app.l = lg.Get()

	return app
}

func (a *App) Run() {
	go func() {
		if err := a.httpServer.ListenAndServe(); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				lg.Fatal(err.Error())
			}
		}
	}()

	lg.Infof("Started on %s", a.cfg.HTTPServer.Addr)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	<-ctx.Done()
	lg.Info("Shutdown success")
}

func (a *App) initHTTPServer() *http.Server {
	mux := http.NewServeMux()

	server := &http.Server{
		Addr:         a.cfg.HTTPServer.Addr,
		Handler:      mux,
		IdleTimeout:  15 * time.Second,
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
	}

	if err := api.MapHandlers(mux, a.impl.MetricService); err != nil {
		lg.Fatalf("unable to route mux: %v", err)
	}

	return server
}
