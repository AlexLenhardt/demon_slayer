package handlers

import (
	"net/http"
	"sample4doc_go/service"
)

// ProdutosHandler utiliza o serviço de produtos para implementar os handlers
type ProdutosHandler struct {
	produtoService service.ProdutoService
}

func NewProdutosHandler(produtoService service.ProdutoService) *ProdutosHandler {
	return &ProdutosHandler{produtoService: produtoService}
}

func (handler *ProdutosHandler) ListarProdutos(w http.ResponseWriter, r *http.Request) {
	handler.produtoService.ListarProdutos(w, r)
}

// ... (Implementar handlers para os demais endpoints)
