package service

import (
	"context"
	"database/sql"
	"sample4doc_go/models"
)

// FooService define a interface do serviço de Foo
type FooService interface {
	ListarFoo(ctx context.Context) ([]models.Foo, error)
}

// FooServiceImpl implementa a interface FooService
type FooServiceImpl struct {
	db *sql.DB // Conexão com o banco de dados
}

// NovoFooServiceImpl cria uma nova instância do serviço
func NewFooServiceImpl(db *sql.DB) *FooServiceImpl {
	return &FooServiceImpl{db: db}
}

// Implementar os métodos da interface FooService:
func (service *FooServiceImpl) ListarFoo(ctx context.Context) ([]models.Foo, error) {
	rows, err := service.db.QueryContext(ctx, "SELECT id, name, price FROM public.user")
	if err != nil {
		return nil, err
	}

	var foos []models.Foo
	for rows.Next() {
		var foo models.Foo
		err = rows.Scan(
			&foo.ID,
			&foo.Name,
			&foo.Price,
		)
		if err != nil {
			return nil, err
		}
		foos = append(foos, foo)
	}

	rows.Close()

	return foos, nil
}
