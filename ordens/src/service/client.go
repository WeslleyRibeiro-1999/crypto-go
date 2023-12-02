package service

import (
	"context"
	"log"

	"github.com/WeslleyRibeiro-1999/crypto-go/ordens/proto/pb"
	"google.golang.org/grpc"
)

var userServiceClient pb.UserServiceClient

func InitializeClient(conn *grpc.ClientConn) {
	userServiceClient = pb.NewUserServiceClient(conn)
}

func GetUserFromGRPC(userID int64) (*pb.GetUserResponse, error) {
	response, err := userServiceClient.GeUserMessage(context.Background(), &pb.GetUserRequest{UserID: userID})
	if err != nil {
		log.Println("erro ao chamar GeUserMessage()")
		return nil, err
	}

	return response, nil
}
