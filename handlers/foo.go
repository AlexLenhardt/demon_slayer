package handlers

import (
	"encoding/json"
	"net/http"
	"sample4doc_go/service"
)

// FooHandler utiliza o serviço de Foos para implementar os handlers
type FooHandler struct {
	fooService service.FooService
}

func NewFooHandler(fooService service.FooService) *FooHandler {
	return &FooHandler{fooService: fooService}
}

func (handler *FooHandler) ListarFoo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := r.Context()
	Foos, err := handler.fooService.ListarFoo(ctx)
	if err != nil {
		http.Error(w, err.Error(), 404)
	}

	buf, _ := json.Marshal(Foos)

	w.Write([]byte(buf))
}
