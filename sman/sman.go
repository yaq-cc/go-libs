package sman

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type WebServiceManager struct {
	Context  context.Context
	Server   *http.Server
	ServeMux *http.ServeMux
	Logger   *log.Logger
	Signals  chan os.Signal
}

func NewWebServiceManager() *WebServiceManager {
	ctx := context.Background()
	mux := http.NewServeMux()
	logger := log.Default()
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)
	return &WebServiceManager{
		Context:  ctx,
		ServeMux: mux,
		Logger:   logger,
		Signals:  signals,
	}
}

func (m *WebServiceManager) HandleFunc(pattern string, handler func(w http.ResponseWriter, r *http.Request)) {
	m.ServeMux.HandleFunc(pattern, handler)
}

func (m *WebServiceManager) ListenAndServe(addr string) error {
	server := &http.Server{
		Addr:    addr,
		Handler: m.ServeMux,
	}
	m.Server = server // Not necessary- done for consistency.
	m.Logger.Printf("Starting HTTP(S) server on port %s\n", addr)
	errs := make(chan error)
	go func() {
		err := m.Server.ListenAndServe()
		if err != nil {
			errs <- err
		}
	}()
Listener:
	for {
		select {
		case err:= <-errs:
			close(errs)
			return err
		case sig := <-m.Signals:
			switch sig {
			case syscall.SIGHUP:
				// pass
			case syscall.SIGINT, syscall.SIGTERM:
				shutdown, cancel := context.WithTimeout(m.Context, time.Second*5)
				defer cancel()
				m.Logger.Println("Gracefully shutting HTTP(S) server down")
				m.Server.Shutdown(shutdown)
				close(m.Signals)
				break Listener
			}
		}
	}
	return nil
}

func (m *WebServiceManager) GetLogger() *log.Logger {
	return m.Logger
}
