package server

import (
	"context"
	"ipfs-file-api/internal/config"
	"ipfs-file-api/pkg/graceful"
	"log/slog"
	"net/http"
	"strconv"
	"time"
)

type Server struct {
	*http.Server
}

type OptionFunc func(*Server)

// new HTTP server
func NewServer(handler http.Handler, optionFuncs ...OptionFunc) *Server {
	srv := &Server{
		Server: &http.Server{
			Addr:    ":" + strconv.Itoa(config.ApiPort),
			Handler: handler,
		},
	}

	for _, optionFs := range optionFuncs {
		optionFs(srv)
	}

	return srv
}

// run HTTP server
func (srv *Server) Start() {
	mgr := graceful.GetManager()

	mgr.Go(func(ctx context.Context) error {
		slog.Info("Running HTTP server", "address", srv.Addr)
		if err := srv.Server.ListenAndServe(); err != nil {
			slog.Error(err.Error())
			return err
		}

		return nil
	})

	mgr.RegisterOnShutdown(func() error {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			slog.Error("Shutdown HTTP server", "err", err)
			return err
		}

		return nil
	})
}
