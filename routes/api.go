package routes

import (
	"goapi/handlers"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func InitRoutes(db *gorm.DB) *mux.Router {
	r := mux.NewRouter()
	userRouter := r.PathPrefix("/users").Subrouter()

    // GET routes
	userRouter.HandleFunc("", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetUser(w, r)
	}).Methods(http.MethodGet)

	userRouter.HandleFunc("/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetUser(w, r)
	}).Methods(http.MethodGet)

    // POST routes 

	userRouter.HandleFunc("", func(w http.ResponseWriter, r *http.Request) {
		handlers.CreateUser(db, w, r)
	}).Methods(http.MethodPost)

    // PUT routes

	return r
}
