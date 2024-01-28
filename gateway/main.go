package main

import (
	"context"
	"flag"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	friendsv1 "github.com/rose-m/monoone/gen/proto/go/friends/v1"
	"github.com/rs/cors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	// env flags:
	// development mode
	devMode = flag.Bool("dev", false, "development mode")
	// command-line options:
	// gRPC server endpoint
	grpcServerEndpoint = flag.String("apiserver-endpoint", "localhost:9090", "gRPC apiserver server endpoint")
)

func main() {
	flag.Parse()

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := friendsv1.RegisterFriendsServiceHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
	if err != nil {
		log.Fatalf("failed to register apiserver gRPC handler: %v", err)
	}

	var handler http.Handler
	handler = mux
	if *devMode {
		handler = cors.New(cors.Options{
			AllowedOrigins:   []string{"http://localhost:5173"},
			AllowedMethods:   []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"ACCEPT", "Authorization", "Content-Type", "X-CSRF-Token"},
			AllowCredentials: true,
		}).Handler(mux)
	}

	server := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	log.Printf("Starting gRPC gateway on port 8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("failed to serve gRPC gateway: %v", err)
	}
}
