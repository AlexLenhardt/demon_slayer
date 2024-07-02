package service

import (
	"database/sql"
	"net/http"
)

// ProdutoService define a interface do serviço de produtos
type ProdutoService interface {
	ListarProdutos(w http.ResponseWriter, r *http.Request)
}

// ProdutosServiceImpl implementa a interface ProdutoService
type ProdutosServiceImpl struct {
	db *sql.DB // Conexão com o banco de dados
}

// NovoProdutosServiceImpl cria uma nova instância do serviço
func NewProdutosServiceImpl(db *sql.DB) *ProdutosServiceImpl {
	return &ProdutosServiceImpl{db: db}
}

// Implementar os métodos da interface ProdutoService:
func (service *ProdutosServiceImpl) ListarProdutos(w http.ResponseWriter, r *http.Request) {
	// ... (Implementação da lógica para listar produtos)
}
