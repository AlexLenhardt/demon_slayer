package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sample4doc_go/models"
	"sample4doc_go/service"

	"github.com/gorilla/mux"
	"github.com/markbates/goth/gothic"
)

// UserHandler utiliza o serviço de Users para implementar os handlers
type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) {
	handler := UserHandler{userService: userService}

	router := mux.NewRouter()
	router.PathPrefix("/user")
	router.HandleFunc("/", handler.PostUser).Methods("POST")
}
func (handler *UserHandler) PostUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := r.Context()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte(err.Error()))
	}

	var jsonUser *models.User
	err = json.Unmarshal(body, jsonUser)
	if err != nil {
		w.Write([]byte(err.Error()))
	}

	if jsonUser != nil {
		user, err := handler.userService.CreateUser(ctx, *jsonUser)
		if err != nil {
			w.Write([]byte(err.Error()))
		}

		b, err := json.Marshal(user)
		if err != nil {
			w.Write([]byte(err.Error()))
		}

		w.Write(b)
	}
}

func (handler *UserHandler) GetAuthCallbackFunction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		fmt.Fprint(w, r)
		return
	}

	fmt.Println(user)

	http.Redirect(w, r, "http://localhost:8080", http.StatusFound)
}
