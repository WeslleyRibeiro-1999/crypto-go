package main

import (
	"fmt"

	"github.com/WeslleyRibeiro-1999/crypto-go/users/db"
	"github.com/WeslleyRibeiro-1999/crypto-go/users/pkg/api"
	"github.com/WeslleyRibeiro-1999/crypto-go/users/pkg/repository"
	"github.com/WeslleyRibeiro-1999/crypto-go/users/pkg/usecase"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
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

	e := echo.New()

	e.POST("/user", httpUser.CreateUser)
	e.GET("/user/:id", httpUser.GetUserID)
	e.GET("/users", httpUser.GetAllUsers)
	e.PUT("/user", httpUser.UpdateUser)
	e.DELETE("/user/:id", httpUser.DeleteUser)
	e.Logger.Fatal(e.Start(":8080"))
}
