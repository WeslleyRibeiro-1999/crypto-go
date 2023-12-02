package repository

import (
	"testing"

	"github.com/WeslleyRibeiro-1999/crypto-go/ordens/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Erro ao criar banco de dados em memória: %v", err)
	}

	err = db.AutoMigrate(&models.Order{})
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

func TestCreateOrder(t *testing.T) {
	db := setupTestDB(t)
	repository := NewRepository(db)
	defer closeTestDB(t, db)

	tests := []struct {
		name string
		args *models.Order
		want *models.Order
	}{
		{
			name: "devem ser iguai",
			args: &models.Order{
				ID:        1,
				UserID:    1,
				Pair:      "1234",
				Amount:    12.33,
				Direction: "sasadafnjdfd",
				Type:      "market",
			},
			want: &models.Order{
				ID:     1,
				UserID: 1,
				Amount: 12.33,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			order, err := repository.CreateOrder(tt.args)
			assert.NoError(t, err, "erro ao criar pedido")

			assert.Equal(t, tt.want.ID, order.ID, "IDs devem ser iguai")
			assert.Equal(t, tt.want.UserID, order.UserID, "user id devem ser iguai")
			assert.Equal(t, tt.want.Amount, order.Amount, "amount devem ser iguai")
		})
	}
}

func TestFindAllByUserID(t *testing.T) {
	db := setupTestDB(t)
	repository := NewRepository(db)
	defer closeTestDB(t, db)

	tests := []struct {
		name string
		args *models.Order
		want *models.Order
	}{
		{
			name: "devem ser iguai",
			args: &models.Order{
				ID:        1,
				UserID:    1,
				Pair:      "1234",
				Amount:    12.33,
				Direction: "sasadafnjdfd",
				Type:      "market",
			},
			want: &models.Order{
				ID:     1,
				UserID: 1,
				Amount: 12.33,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := repository.CreateOrder(tt.args)
			assert.NoError(t, err, "erro ao criar pedido")

			order, err := repository.FindAllByUserID(tt.args.UserID)
			assert.NoError(t, err, "erro ao buscar pedido")

			for _, v := range *order {
				assert.Equal(t, tt.want.ID, v.ID, "IDs devem ser iguai")
				assert.Equal(t, tt.want.UserID, v.UserID, "user id devem ser iguai")
				assert.Equal(t, tt.want.Amount, v.Amount, "amount devem ser iguai")
			}

		})
	}
}

func TestFindOneByID(t *testing.T) {
	db := setupTestDB(t)
	repository := NewRepository(db)
	defer closeTestDB(t, db)

	tests := []struct {
		name string
		args *models.Order
		want *models.Order
	}{
		{
			name: "devem ser iguai",
			args: &models.Order{
				ID:        1,
				UserID:    1,
				Pair:      "1234",
				Amount:    12.33,
				Direction: "sasadafnjdfd",
				Type:      "market",
			},
			want: &models.Order{
				ID:     1,
				UserID: 1,
				Amount: 12.33,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := repository.CreateOrder(tt.args)
			assert.NoError(t, err, "erro ao criar pedido")

			order, err := repository.FindOneByID(tt.args.UserID)
			assert.NoError(t, err, "erro ao buscar pedido")

			assert.Equal(t, tt.want.ID, order.ID, "IDs devem ser iguai")
			assert.Equal(t, tt.want.UserID, order.UserID, "user id devem ser iguai")
			assert.Equal(t, tt.want.Amount, order.Amount, "amount devem ser iguai")
		})
	}
}

func TestFindAllOrders(t *testing.T) {
	db := setupTestDB(t)
	repository := NewRepository(db)
	defer closeTestDB(t, db)

	tests := []struct {
		name string
		args *models.Order
		want *models.Order
	}{
		{
			name: "devem ser iguai",
			args: &models.Order{
				ID:        1,
				UserID:    1,
				Pair:      "1234",
				Amount:    12.33,
				Direction: "sasadafnjdfd",
				Type:      "market",
			},
			want: &models.Order{
				ID:     1,
				UserID: 1,
				Amount: 12.33,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := repository.CreateOrder(tt.args)
			assert.NoError(t, err, "erro ao criar pedido")

			order, err := repository.FindAllOrders()
			assert.NoError(t, err, "erro ao buscar pedidos")

			for _, v := range *order {
				assert.Equal(t, tt.want.ID, v.ID, "IDs devem ser iguai")
				assert.Equal(t, tt.want.UserID, v.UserID, "user id devem ser iguai")
				assert.Equal(t, tt.want.Amount, v.Amount, "amount devem ser iguai")
			}

		})
	}
}

func TestDeleteUser(t *testing.T) {
	db := setupTestDB(t)
	repository := NewRepository(db)
	defer closeTestDB(t, db)

	tests := []struct {
		name string
		args *models.Order
		want bool
	}{
		{
			name: "devem ser iguai",
			args: &models.Order{
				ID:        1,
				UserID:    1,
				Pair:      "1234",
				Amount:    12.33,
				Direction: "sasadafnjdfd",
				Type:      "market",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := repository.CreateOrder(tt.args)
			assert.NoError(t, err, "erro ao criar pedido")

			err = repository.DeleteOrder(tt.args.ID)
			assert.NoError(t, err, "erro ao buscar pedido")
		})
	}
}
