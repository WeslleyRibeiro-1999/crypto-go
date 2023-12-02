package main

import (
	"fmt"
	"log"

	"github.com/WeslleyRibeiro-1999/crypto-go/ordens/db"
	"github.com/WeslleyRibeiro-1999/crypto-go/ordens/src/api"
	"github.com/WeslleyRibeiro-1999/crypto-go/ordens/src/repository"
	"github.com/WeslleyRibeiro-1999/crypto-go/ordens/src/service"
	"github.com/WeslleyRibeiro-1999/crypto-go/ordens/src/usecase"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
	http := api.NewHandler(usecase)

	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("nao foi possivel se conectar: %v", err)
	}
	defer conn.Close()

	service.InitializeClient(conn)

	e := echo.New()

	e.POST("/order", http.CreateOrder)
	e.GET("/orders/:user_id", http.GetAllbyUserID)
	e.GET("/orders", http.GetAllOrders)
	e.GET("/order/:id", http.GetOrderByID)
	e.DELETE("order/:id", http.DeleteOrder)
	e.Logger.Fatal(e.Start(":8000"))
}
