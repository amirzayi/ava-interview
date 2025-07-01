package test

import (
	"context"
	"log"
	"testing"

	"github.com/amirzayi/ava-interview/app"
)

func TestMain(m *testing.M) {
	log.Print("setting up")

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		defer cancel()
		err := app.Start(ctx, "file:inmem_db?mode=memory&cache=shared", ":8080")
		if err != nil {
			log.Fatal(err)
		}
	}()

	m.Run()
}
