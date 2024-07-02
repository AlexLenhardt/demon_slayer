package main

import (
	"fmt"
	"net/http"
	"sample4doc_go/db"
	"sample4doc_go/handlers"
	"sample4doc_go/service"

	"github.com/gorilla/mux"
)

func main() {
	// Abre a conexão com o banco de dados usando as configurações carregadas
	db, err := db.ConectarDB()
	if err != nil {
		panic(err)
	}

	// Criar instância do serviço de products
	productService := service.NewProductsServiceImpl(db)

	// Criar instância do handler de products
	productsHandler := handlers.NewProductsHandler(productService)

	// Criar roteador e registrar handlers
	router := mux.NewRouter()
	router.HandleFunc("/products", productsHandler.ListarProducts).Methods("GET")

	// ... (Registrar outros handlers para os demais endpoints)

	// Iniciar servidor HTTP
	fmt.Println("Servidor iniciado na porta 8080")
	http.ListenAndServe(":8080", router)
}
