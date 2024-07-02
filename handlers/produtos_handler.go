package handlers

import (
	"encoding/json"
	"net/http"
	"sample4doc_go/service"
)

// ProductsHandler utiliza o serviço de products para implementar os handlers
type ProductsHandler struct {
	productService service.ProductService
}

func NewProductsHandler(productService service.ProductService) *ProductsHandler {
	return &ProductsHandler{productService: productService}
}

func (handler *ProductsHandler) ListarProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := r.Context()
	products, err := handler.productService.ListarProducts(ctx)
	if err != nil {
		http.Error(w, err.Error(), 404)
	}

	buf, _ := json.Marshal(products)

	w.Write([]byte(buf))
}
