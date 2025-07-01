package app

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	_ "github.com/mattn/go-sqlite3"

	"github.com/amirzayi/ava-interview/api"
	"github.com/amirzayi/ava-interview/service"
)

func Start(ctx context.Context, dbPath, httpAddress string) error {
	// Initialize the database connection
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return fmt.Errorf("failed to open database connection %v", err)
	}
	err = db.PingContext(context.Background())
	if err != nil {
		return fmt.Errorf("failed to ping database %v", err)
	}
	log.Print("db ok")
	// Initialize the service(business logics)
	service := service.NewService(db)

	// Initialize the router(restful apis)
	router := api.Router{
		Service: service,
	}

	// Start the http server
	mux := http.NewServeMux()
	router.Register(mux)

	srv := &http.Server{
		Addr:    httpAddress,
		Handler: mux,
	}

	errCh := make(chan error)

	exitCtx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	go func() {
		err = srv.ListenAndServe()
		if !errors.Is(err, http.ErrServerClosed) {
			errCh <- fmt.Errorf("failed to start http server %v", err)
		}
	}()
	log.Printf("http server started on %s\n", srv.Addr)

	select {
	case err = <-errCh:
		return err

	case <-exitCtx.Done():
		log.Println("received terminate signal")
	}

	// Gracefully Close the http server and database connection
	return errors.Join(srv.Shutdown(ctx), srv.Close(), db.Close())
}
