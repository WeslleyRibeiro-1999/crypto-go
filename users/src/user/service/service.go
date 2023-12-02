package service

import (
	"context"

	"github.com/WeslleyRibeiro-1999/crypto-go/users/proto/pb"
	"github.com/WeslleyRibeiro-1999/crypto-go/users/src/user/repository"
)

type Service struct {
	repo repository.Repository
	pb.UnimplementedUserServiceServer
}

func NewService(repo repository.Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) GeUserMessage(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	userID := req.GetUserID()

	user, err := s.repo.GetUserID(userID)
	if err != nil {
		return nil, err
	}

	res := &pb.GetUserResponse{
		Name: user.Name,
	}

	return res, nil
}
