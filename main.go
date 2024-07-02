package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"sample4doc_go/db"
	"sample4doc_go/handlers"
	"sample4doc_go/service"

	"github.com/gorilla/mux"
)

func main() {
	// Carrega as configurações do banco de dados
	config, err := db.LoadConfig()
	if err != nil {
		panic(err)
	}

	// Abre a conexão com o banco de dados usando as configurações carregadas
	db, err := sql.Open("postgres", config.ToDBURL())
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Criar instância do serviço de produtos
	produtoService := service.NewProdutosServiceImpl(db)

	// Criar instância do handler de produtos
	produtosHandler := handlers.NewProdutosHandler(produtoService)

	// Criar roteador e registrar handlers
	router := mux.NewRouter()
	router.HandleFunc("/produtos", produtosHandler.ListarProdutos).Methods("GET")
	// ... (Registrar outros handlers para os demais endpoints)

	// Iniciar servidor HTTP
	fmt.Println("Servidor iniciado na porta 8080")
	http.ListenAndServe(":8080", router)
}
