package main

import (
	"context"
	"errors"
	"github.com/go-chi/chi"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const addr = ":8080"

func main() {
	ctx := context.Background()

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		select {
		case <-ctx.Done():
			return ctx.Err()
		case sig := <-c:
			return errors.New("graceful shutdown by " + sig.String())
		}
	})

	server := &http.Server{
		Addr:    addr,
		Handler: chi.ServerBaseContext(ctx, NewHandler()),
	}
	g.Go(func() error {
		<-ctx.Done()
		return server.Shutdown(context.Background())
	})
	g.Go(func() error {
		log.Println("start listening at", addr, "...")
		return server.ListenAndServe()
	})

	err := g.Wait()
	log.Println("server stopped:", err)
}

func NewHandler() http.Handler {
	router := chi.NewRouter()
	router.Get("/sleep", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		t := time.Second

		ts := r.URL.Query().Get("time")
		if ts != "" {
			var err error
			t, err = time.ParseDuration(ts)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Header().Set("Content-Type", "text/plain")
				_, _ = w.Write([]byte(err.Error()))
				return
			}
		}

		timer := time.NewTimer(t)
		select {
		case <-timer.C:
		case <-ctx.Done():
			timer.Stop()
			log.Println("interrupted by context")
			w.WriteHeader(http.StatusInternalServerError)
			w.Header().Set("Content-Type", "text/plain")
			_, _ = w.Write([]byte(ctx.Err().Error()))
			return
		}

		_, _ = w.Write([]byte("Hello after " + t.String()))
	})
	return router
}
