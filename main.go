package main

import (
	"fmt"
	"os"
	"time"

	"github.com/joaovictorsl/proto-sample/proto/user"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	birthdate := time.Date(2003, time.February, 20, 0, 0, 0, 0, time.UTC)
	me := &user.User{
		Name:      "João Victor",
		Password:  "senhaboa",
		Type:      user.UserType_PERSONAL,
		Birthdate: timestamppb.New(birthdate),
	}

	userToFile(me, "./me")
	u := userFromFile("./me")

	d, _ := os.ReadFile("./me")
	fmt.Println(d)
	// Name
	fmt.Println(string(d[2 : 2+len(me.Name)]))
	// Password
	fmt.Println(string(d[2+len(me.Name)+2 : 2+len(me.Name)+2+len(me.Password)]))
	// [10 12 J o a ~ o   V i c t o r 18 8 115 101 110 104 97 98 111 97 34 6 8 128 181 208 242 3]

	fmt.Println(me.String())
	fmt.Println(u.String())
}

func userToFile(data *user.User, filename string) {
	raw, err := proto.Marshal(data)
	if err != nil {
		panic(err)
	}

	if err := os.WriteFile(filename, raw, 0777); err != nil {
		panic(err)
	}
}

func userFromFile(filename string) *user.User {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var user user.User
	err = proto.Unmarshal(data, &user)
	if err != nil {
		panic(err)
	}

	return &user
}
