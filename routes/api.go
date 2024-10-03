package routes

import (
	"goapi/handlers"
	"goapi/middleware"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func InitRoutes(db *gorm.DB) *mux.Router {
	r := mux.NewRouter()

	userRouter := r.PathPrefix("/users").Subrouter()
	protectedRouter := r.PathPrefix("/protected").Subrouter()
	protectedRouter.Use(middleware.JWTMiddleware)

	// GET routes public
	userRouter.HandleFunc("", handlers.GetUser).Methods(http.MethodGet)
	userRouter.HandleFunc("/{id}", handlers.GetUser).Methods(http.MethodGet)

	// GET routes private
	protectedRouter.HandleFunc("/user/{id}", middleware.ProtectedHandler).Methods(http.MethodGet)

	// POST routes public
	userRouter.HandleFunc("",
		func(w http.ResponseWriter, r *http.Request) {
			handlers.CreateUser(db, w, r)
		}).Methods(http.MethodPost)

	r.HandleFunc("/login", handlers.LoginHandler).Methods(http.MethodPost)

	// POST routes private

	// PUT routes
	protectedRouter.HandleFunc("/user/{id}",
		func(w http.ResponseWriter, r *http.Request) {
			handlers.UpdateUser(db, w, r)
		}).Methods(http.MethodPut)

	// DELETE routes
	protectedRouter.HandleFunc("/user/{id}",
		func(w http.ResponseWriter, r *http.Request) {
			handlers.DeleteUser(db, w, r)
		}).Methods(http.MethodDelete)

	return r
}
