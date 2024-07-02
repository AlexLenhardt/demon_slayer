package service

import (
	"context"
	"database/sql"
	"sample4doc_go/models"
)

// ProductService define a interface do serviço de products
type ProductService interface {
	ListarProducts(ctx context.Context) ([]models.Product, error)
}

// ProductsServiceImpl implementa a interface ProductService
type ProductsServiceImpl struct {
	db *sql.DB // Conexão com o banco de dados
}

// NovoProductsServiceImpl cria uma nova instância do serviço
func NewProductsServiceImpl(db *sql.DB) *ProductsServiceImpl {
	return &ProductsServiceImpl{db: db}
}

// Implementar os métodos da interface ProductService:
func (service *ProductsServiceImpl) ListarProducts(ctx context.Context) ([]models.Product, error) {
	rows, err := service.db.QueryContext(ctx, "SELECT id, name, price FROM public.user")
	if err != nil {
		return nil, err
	}

	var products []models.Product
	for rows.Next() {
		var product models.Product
		err = rows.Scan(
			&product.ID,
			&product.Name,
			&product.Price,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	rows.Close()

	return products, nil
}
