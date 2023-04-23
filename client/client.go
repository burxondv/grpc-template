package main

import (
	"context"
	"log"
	"time"

	pb "github.com/grpc-template/protos"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:9090", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Println("failed to connect: ", err)
	}
	defer conn.Close()

	c := pb.NewUserManagementClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var new_user = make(map[string]int32)
	new_user["Alice"] = 43
	new_user["Bob"] = 30
	for name, age := range new_user {
		r, err := c.CreateNewUser(ctx, &pb.UserRequest{Name: name, Age: age})
		if err != nil {
			log.Printf("failed to create user: %v", err)
		}
		log.Printf(`User details:
		NAME: %s
		AGE: %d
		ID: %d`, r.GetName(), r.GetAge(), r.GetId())
	}

}
