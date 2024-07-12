package main

import (
	"fmt"
	"net/http"
	"sample4doc_go/auth"
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

	auth.NewAuth()

	// Criar instância do serviço de products
	fooService := service.NewFooServiceImpl(db)
	userService := service.NewUserServiceImpl(db)

	// Criar instância do handler de foos
	fooHandler := handlers.NewFooHandler(fooService)
	handlers.NewUserHandler(userService)

	// Criar roteador e registrar handlers
	router := mux.NewRouter()
	router.HandleFunc("/foo", fooHandler.ListarFoo).Methods("GET")

	// ... (Registrar outros handlers para os demais endpoints)

	// Iniciar servidor HTTP
	fmt.Println("Servidor iniciado na porta 8080")
	http.ListenAndServe(":8080", router)
}
