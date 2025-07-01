package test

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/amirzayi/ava-interview/app"
)

func TestMain(m *testing.M) {

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		defer cancel()
		err := app.Start(ctx, "file:inmem_db?mode=memory&cache=shared", ":8080")
		if err != nil {
			log.Fatal(err)
		}
	}()
	// wait a while for the server to start
	time.Sleep(500 * time.Millisecond)
	m.Run()
}
