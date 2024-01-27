package main

import (
	"fmt"

	"github.com/google/uuid"
	friendsv1 "github.com/rose-m/monoone/gen/proto/go/friends/v1"
)

func main() {
	fmt.Printf("Hello, World: %s\n", uuid.NewString())

	request := &friendsv1.GetFriendsRequest{}
	fmt.Printf("request: %v\n", request)
}
