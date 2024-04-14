package main

import (
	"context"
	"fmt"

	"log"

	"github.com/shohinsan/GopherQL/config"
	"github.com/shohinsan/GopherQL/postgres"
)

func main() {
	ctx := context.Background()

	conf := config.New()

	db := postgres.New(ctx, conf)

	if err := db.Migrate(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("WORKING")
}
