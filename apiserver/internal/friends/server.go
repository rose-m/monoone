package friends

import (
	"context"
	"log"

	friendsv1 "github.com/rose-m/monoone/gen/proto/go/friends/v1"
)

type Server struct {
	friendsv1.UnimplementedFriendsServiceServer
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) GetFriends(ctx context.Context, request *friendsv1.GetFriendsRequest) (*friendsv1.GetFriendsResponse, error) {
	log.Printf("Received request: %v", request)
	return &friendsv1.GetFriendsResponse{}, nil
}
