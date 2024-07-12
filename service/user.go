package service

import (
	"context"
	"database/sql"
	"log"
	"sample4doc_go/models"
)

// UserService define a interface do serviço de User
type UserService interface {
	ListUser(ctx context.Context) ([]models.User, error)
	CreateUser(ctx context.Context, user models.User) (*models.User, error)
}

// UserServiceImpl implementa a interface UserService
type UserServiceImpl struct {
	db *sql.DB // Conexão com o banco de dados
}

// NovoUserServiceImpl cria uma nova instância do serviço
func NewUserServiceImpl(db *sql.DB) *UserServiceImpl {
	return &UserServiceImpl{db: db}
}

func (service *UserServiceImpl) CreateUser(ctx context.Context, user models.User) (*models.User, error) {
	queryInsert := `
	INSERT INTO public."user"(name, birth) VALUES (?, ?)`

	res, err := service.db.ExecContext(ctx, queryInsert, user.Name, user.Birth)
	if err != nil {
		log.Fatal("Error during ExecContext CreateUser", err.Error())
		return nil, err
	}

	user.ID, err = res.LastInsertId()
	if err != nil {
		log.Fatal("Error during LastInsertId CreateUser", err.Error())
		return nil, err
	}

	return &user, nil
}

// Implementar os métodos da interface UserService:
func (service *UserServiceImpl) ListUser(ctx context.Context) ([]models.User, error) {
	rows, err := service.db.QueryContext(ctx, "SELECT id, name FROM public.user")
	if err != nil {
		return nil, err
	}

	var users []models.User
	for rows.Next() {
		var user models.User
		err = rows.Scan(
			&user.ID,
			&user.Name,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	rows.Close()

	return users, nil
}
