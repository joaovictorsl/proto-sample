package svc

import (
	"context"
	"sync"

	"github.com/joaovictorsl/proto-sample/proto/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var storage = make(map[int32]*user.User)
var currId = 0

type UserServiceServer struct {
	user.UnimplementedUserServiceServer

	rwLock sync.RWMutex
}

func NewUserServiceServer() *UserServiceServer {
	return &UserServiceServer{
		rwLock: sync.RWMutex{},
	}
}

func (s *UserServiceServer) CreateUser(ctx context.Context, req *user.CreateUserRequest) (res *user.CreateUserResponse, err error) {
	s.rwLock.Lock()
	defer s.rwLock.Unlock()

	if req.Name == "" {
		err = status.New(codes.InvalidArgument, "Name cannot be empty").Err()
		return nil, err
	} else if len(req.Password) < 3 {
		err = status.New(codes.InvalidArgument, "Password length should be >= 3").Err()
		return nil, err
	}

	storage[int32(currId)] = &user.User{
		Id:        int32(currId),
		Name:      req.Name,
		Password:  req.Password,
		Type:      req.Type,
		Birthdate: req.Birthdate,
		Friends:   req.Friends,
	}

	res = &user.CreateUserResponse{
		UserId: int32(currId),
	}

	currId++

	return res, nil
}

func (s *UserServiceServer) GetUser(ctx context.Context, req *user.GetUserRequest) (res *user.GetUserResponse, err error) {
	s.rwLock.RLock()
	defer s.rwLock.RUnlock()

	u, ok := storage[req.UserId]
	if !ok {
		return nil, status.New(codes.NotFound, "User not found").Err()
	}

	res = &user.GetUserResponse{
		User: u,
	}

	return res, nil
}
