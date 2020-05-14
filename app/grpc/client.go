package grpc

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/vsokoltsov/users-service/app/grpc/authorization"
	"google.golang.org/grpc"
)

const (
	address = "authorization360:50001"
)

// GetUserId sends token to authorization service
// and receives user's id if it is present
func GetUserId(token string) (*int, error) {
	var userId int
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(5*time.Second))
	if err != nil {
		// log.Fatalf("did not connect: %v", err)
		return nil, errors.New(fmt.Sprintf("Cannot connect to the authotization service: %s ", err.Error()))
	}
	c := authorization.NewAuthorizationServiceClient(conn)
	// var deadlineMs = flag.Int("deadline_ms", 20*1000, "Default deadline in milliseconds.")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	r, err := c.GetUserId(ctx, &authorization.TokenRequest{Token: token})
	if err != nil {
		// log.Fatalf("could not greet: %v", err)
		return nil, err
	}
	log.Printf("THE ANSWER IS %s", r.GetUserID())
	userId = int(r.GetUserID())
	return &userId, nil
}
