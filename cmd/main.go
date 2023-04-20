package main

import (
	"context"
	"log"
	"transactions/config"
	"transactions/internal/service"
	"transactions/internal/storage"
	"transactions/internal/transport"
)

func main() {
	log.Fatal(run())
}
func run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	conf, err := config.New()
	if err != nil {
		return err
	}
	repo, err := storage.New(ctx, conf)
	if err != nil {
		return err
	}
	svc, err := service.NewManager(repo)
	if err != nil {
		return err
	}
	err = transport.Run(svc)
	if err != nil {
		return err
	}
	return nil
}
