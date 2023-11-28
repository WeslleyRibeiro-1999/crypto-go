package repository

import (
	"fmt"
	"testing"

	"github.com/WeslleyRibeiro-1999/crypto-go/users/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Erro ao criar banco de dados em memória: %v", err)
	}

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		t.Fatalf("Erro ao executar migrações: %v", err)
	}

	return db
}

func closeTestDB(t *testing.T, db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		t.Fatalf("erro ao fechar conexao banco de dados em memória: %v", err)
	}
	sqlDB.Close()
}

func TestCreateUser(t *testing.T) {
	db := setupTestDB(t)
	repository := NewRepository(db)
	defer closeTestDB(t, db)

	tests := []struct {
		name string
		args *models.User
		want *models.User
	}{
		{
			name: "devem ser iguais",
			args: &models.User{
				ID:          1,
				Name:        "weslley",
				Email:       "teste@teste.com",
				PhoneNumber: "1199999999",
			},
			want: &models.User{
				ID:          1,
				Name:        "weslley",
				Email:       "teste@teste.com",
				PhoneNumber: "1199999999",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user, err := repository.CreateUser(tt.args)
			assert.NoError(t, err, "erro ao criar usuario")

			assert.Equal(t, tt.want.Email, user.Email, "emails devem ser iguais")
			assert.Equal(t, tt.want.Name, user.Name, "nomes devem ser iguais")
			assert.Equal(t, tt.want.ID, user.ID, "IDs devem ser iguais")
		})
	}
}

func TestGetUserID(t *testing.T) {
	db := setupTestDB(t)
	defer closeTestDB(t, db)

	repository := NewRepository(db)

	tests := []struct {
		name string
		args *models.User
		want *models.User
	}{
		{
			name: "devem ser iguais",
			args: &models.User{
				ID:          1,
				Name:        "weslley",
				Email:       "teste@teste.com",
				PhoneNumber: "1199999999",
			},
			want: &models.User{
				ID:    1,
				Name:  "weslley",
				Email: "teste@teste.com",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := repository.CreateUser(tt.args)
			assert.NoError(t, err, "erro ao criar usuario")

			user, err := repository.GetUserID(tt.args.ID)
			assert.NoError(t, err, "erro ao buscar usuario")
			fmt.Println("user: ", user)

			assert.Equal(t, tt.want.Email, user.Email, "emails devem ser iguais")
			assert.Equal(t, tt.want.Name, user.Name, "nomes devem ser iguais")
			assert.Equal(t, tt.want.ID, user.ID, "IDs devem ser iguais")
		})
	}
}

func TestGetAllUsers(t *testing.T) {
	db := setupTestDB(t)
	defer closeTestDB(t, db)

	repository := NewRepository(db)

	tests := []struct {
		name string
		args *[]models.User
		want int
	}{
		{
			name: "deve comparar a quantidade de indices",
			args: &[]models.User{
				{
					ID:          1,
					Name:        "weslley",
					Email:       "teste@teste.com",
					PhoneNumber: "1199999999",
				},
				{
					ID:          2,
					Name:        "luana",
					Email:       "teste1@teste.com",
					PhoneNumber: "1199999999",
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, v := range *tt.args {
				_, err := repository.CreateUser(&v)
				assert.NoError(t, err, "erro ao criar usuario")
			}

			users, err := repository.GetAllUsers()
			assert.NoError(t, err, "erro ao buscar usuarios")
			fmt.Println("user: ", len(*users))

			assert.Equal(t, len(*users), tt.want, "quantidade de indices devem ser iguais")
		})
	}
}

func TestUpdateUser(t *testing.T) {
	db := setupTestDB(t)
	defer closeTestDB(t, db)

	repository := NewRepository(db)

	tests := []struct {
		name       string
		args       *models.User
		updateArgs *models.User
		want       *models.User
	}{
		{
			name: "devem ser iguais",
			args: &models.User{
				ID:          1,
				Name:        "weslley",
				Email:       "teste@teste.com",
				PhoneNumber: "1199999999",
			},
			updateArgs: &models.User{
				ID:          1,
				Name:        "luana",
				Email:       "teste@teste.com",
				PhoneNumber: "1199999999",
			},
			want: &models.User{
				ID:          1,
				Name:        "luana",
				Email:       "teste@teste.com",
				PhoneNumber: "1199999999",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := repository.CreateUser(tt.args)
			assert.NoError(t, err, "erro ao criar usuario")

			user, err := repository.UpdateUser(tt.updateArgs)
			assert.NoError(t, err, "erro ao criar usuario")

			assert.Equal(t, tt.want.Email, user.Email, "emails devem ser iguais")
			assert.Equal(t, tt.want.Name, user.Name, "nomes devem ser iguais")
			assert.Equal(t, tt.want.ID, user.ID, "IDs devem ser iguais")
		})
	}
}

func TestDeleteUser(t *testing.T) {
	db := setupTestDB(t)
	defer closeTestDB(t, db)

	repository := NewRepository(db)

	tests := []struct {
		name     string
		args     *models.User
		want     int64
		expected bool
	}{
		{
			name: "devem ser iguais",
			args: &models.User{
				ID:          1,
				Name:        "weslley",
				Email:       "teste@teste.com",
				PhoneNumber: "1199999999",
			},
			want:     1,
			expected: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := repository.CreateUser(tt.args)
			assert.NoError(t, err, "erro ao criar usuario")

			ok := repository.DeleteUser(tt.want)
			assert.True(t, ok, "erro ao deletar usuario")
		})
	}
}
