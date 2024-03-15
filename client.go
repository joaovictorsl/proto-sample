package main

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/joaovictorsl/proto-sample/proto/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func runClient() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return
	}
	client := user.NewUserServiceClient(conn)

	keepRunning := true
	var input string
	for keepRunning {
		fmt.Println("C to create client, G to get client and E to quit")
		fmt.Scanln(&input)

		switch strings.ToUpper(input) {
		case "C":
			createUser(client)
		case "G":
			getUser(client)
		case "E":
			keepRunning = false
		default:
			fmt.Println("Invalid input, try again")
		}
	}
}

func getUser(client user.UserServiceClient) {
	var idStr string
	fmt.Println("What is the user id?")
	fmt.Scanln(&idStr)

	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Id must be an integer: ", err.Error())
		return
	}

	res, err := client.GetUser(context.Background(), &user.GetUserRequest{UserId: int32(id)})
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res.User)
}

func createUser(client user.UserServiceClient) {
	var name, password string
	fmt.Println("What is the new user name?")
	fmt.Scanln(&name)
	fmt.Println("What is the new user password?")
	fmt.Scanln(&password)

	res, err := client.CreateUser(context.Background(), &user.CreateUserRequest{
		Name:      name,
		Password:  password,
		Type:      user.UserType_PERSONAL,
		Birthdate: timestamppb.New(time.Date(2003, time.February, 20, 13, 30, 0, 0, time.UTC)),
		Friends:   []*user.User{},
	})
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("User ID: %d\n", res.UserId)
}
