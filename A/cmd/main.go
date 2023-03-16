package main

import (
	"A/pb"
	"context"
	"fmt"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"log"
)

func main() {
	// Địa chỉ của server gRPC
	serverAddr := "localhost:9000"

	// Kết nối đến server
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// Khởi tạo client
	client := pb.NewUserServiceClient(conn)
	/*
		// Gọi phương thức List
		fmt.Println("Calling List...")
		reqList := &pb.RequestList{}
		resList, err := client.List(context.Background(), reqList)
		if err != nil {
			log.Fatalf("Error calling List: %v", err)
		}
		fmt.Println("Users:")
		for _, u := range resList.GetUser() {
			fmt.Printf("- %v\n", u)
		}

		// Gọi phương thức FindByEmail
		fmt.Println("Calling FindByEmail...")
		reqFindByEmail := &pb.RequestFindByEmail{Email: "john.doe@example.com"}
		resFindByEmail, err := client.FindByEmail(context.Background(), reqFindByEmail)
		if err != nil {
			log.Fatalf("Error calling FindByEmail: %v", err)
		}
		fmt.Printf("User found: %v\n", resFindByEmail.GetUser())
	*/
	// Gọi phương thức Create
	fmt.Println("Calling Create...")
	reqCreate := &pb.RequestCreate{
		User: &pb.User{
			Id:      uuid.New().String(),
			Name:    "Jane Doe",
			Email:   "jane.doe@example.com",
			Address: "123 Main St",
			Phone:   "555-1234",
		},
	}
	resCreate, err := client.Create(context.Background(), reqCreate)
	if err != nil {
		log.Fatalf("Error calling Create: %v", err)
	}
	fmt.Printf("User created: %v\n", resCreate)
	/*
		// Gọi phương thức UpdateByEmail
		fmt.Println("Calling UpdateByEmail...")
		reqUpdateByEmail := &pb.RequestUpdateByEmail{
			Email: "jane.doe@example.com",
			Fields: &pb.UpdateFields{
				Name:    "Jane Smith",
				Address: "456 Main St",
				Phone:   "555-5678",
			},
		}
		resUpdateByEmail, err := client.UpdateByEmail(context.Background(), reqUpdateByEmail)
		if err != nil {
			log.Fatalf("Error calling UpdateByEmail: %v", err)
		}
		fmt.Printf("Update result: %v\n", resUpdateByEmail.GetMessage())

		// Gọi phương thức DeleteByEmail
		fmt.Println("Calling DeleteByEmail...")
		reqDeleteByEmail := &pb.RequestDeleteByEmail{Email: "jane.doe@example.com"}
		resDeleteByEmail, err := client.DeleteByEmail(context.Background(), reqDeleteByEmail)
		if err != nil {
			log.Fatalf("Error calling DeleteByEmail: %v", err)
		}
		fmt.Printf("Delete result: %v\n", resDeleteByEmail.GetMessage())
	*/
}
