package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/rose-m/monoone/apiserver/internal/friends"
	friendsv1 "github.com/rose-m/monoone/gen/proto/go/friends/v1"
)

func main() {
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	friendsv1.RegisterFriendsServiceServer(s, friends.NewServer())
	log.Printf("Starting server on port 9090")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
