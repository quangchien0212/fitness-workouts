package main

import (
	"context"
	"net/http"
	"os"

	"cloud.google.com/go/firestore"
	"github.com/go-chi/chi/v5"

	"github.com/quangchien0212/fitness-workouts/internal/common/server"
	"github.com/quangchien0212/fitness-workouts/internal/trainings/logs"
	"github.com/quangchien0212/fitness-workouts/internal/trainings/ports"
)

func main() {
	logs.Init()

	ctx := context.Background()
	client, err := firestore.NewClient(ctx, os.Getenv("GCP_PROJECT"))

	if err != nil {
		panic(err)
	}

	defer client.Close()

	server.RunHTTPServer(func(router chi.Router) http.Handler {
		return ports.HandlerFromMux(ports.NewHttpServer(), router)
	})
}
