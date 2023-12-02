package main

import (
	"fmt"
	"log"
	"net"

	"github.com/WeslleyRibeiro-1999/crypto-go/users/db"
	"github.com/WeslleyRibeiro-1999/crypto-go/users/proto/pb"
	"github.com/WeslleyRibeiro-1999/crypto-go/users/src/user/api"
	"github.com/WeslleyRibeiro-1999/crypto-go/users/src/user/repository"
	"github.com/WeslleyRibeiro-1999/crypto-go/users/src/user/service"
	"github.com/WeslleyRibeiro-1999/crypto-go/users/src/user/usecase"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Erro ao carregar arquivo .env: ", err)
		return
	}

	database, err := db.NewDatabase()
	if err != nil {
		panic("Falha ao conectar ao banco de dados")
	}

	repo := repository.NewRepository(database)
	usecase := usecase.NewUsecase(repo)
	httpUser := api.NewHandler(usecase)
	service := service.NewService(repo)

	e := echo.New()

	e.POST("/user", httpUser.CreateUser)
	e.GET("/user/:id", httpUser.GetUserID)
	e.GET("/users", httpUser.GetAllUsers)
	e.PUT("/user", httpUser.UpdateUser)
	e.DELETE("/user/:id", httpUser.DeleteUser)

	go func() {
		e.Logger.Fatal(e.Start(":8080"))
	}()

	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("falha ao subir rpc: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, service)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("falha no servidor: %v", err)
	}
}
