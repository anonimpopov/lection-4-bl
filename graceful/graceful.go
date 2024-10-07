package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

const shutdownTimeout = 15 * time.Second

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/start", StartHandle)

	httpServer := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	group, ctx := errgroup.WithContext(ctx)
	group.Go(func() error {
		if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("err in listen: %s\n", err)
			return fmt.Errorf("failed to serve http server: %w", err)
		}
		fmt.Println("after listener")

		return nil
	})

	group.Go(func() error {
		fmt.Println("before ctx done")
		<-ctx.Done()
		fmt.Println("after ctx done")

		shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
		defer cancel()

		err := httpServer.Shutdown(shutdownCtx)
		if err != nil {
			return err
		}
		fmt.Println("after server shutdown")

		return nil
	})

	err := group.Wait()
	if err != nil {
		fmt.Printf("after wait: %s\n", err)
		return
	}
}

func StartHandle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Start our process"))
	time.Sleep(10 * time.Second)
	fmt.Printf("process was started")
}
