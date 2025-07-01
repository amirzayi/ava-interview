package main

import (
	"context"
	"log"

	"github.com/amirzayi/ava-interview/app"
)

func main() {
	err := app.Start(context.Background(), "./app.db", ":8080")
	if err != nil {
		log.Fatal(err)
	}
}
