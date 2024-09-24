package handlers

import (
	"encoding/json"
	"errors"
	"goapi/dbconnect"
	"goapi/models"
	"net/http"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)


func GetUser(w http.ResponseWriter, r *http.Request) {
	db := dbconnect.DB
    id := mux.Vars(r)["id"]

	// en el caso de que no sea proporcionado un id entra en el if, sino es que se proporciono un id
	if id == "" {
		GetAllUsers(db, w)
	} else {
		GetUserById(db, id, w)
	}

}


func GetUserById(db *gorm.DB, id string, w http.ResponseWriter) {
	var user models.User
	result := db.First(&user, id)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			http.Error(w, "user not found", http.StatusNotFound)
		} else {
			http.Error(w, "error while connecting", http.StatusInternalServerError)
		}
		return
	}

	// mandar una respuesta en json
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)    // decir que la peticion se resolvio sin problemas
	json.NewEncoder(w).Encode(user) // enviar la respuesta en json, la pw no se mostrara
}

func GetAllUsers(db *gorm.DB, w http.ResponseWriter) {
	var users []models.User
	result := db.Find(&users)

	if result.Error != nil {
		http.Error(w, "error while connecting", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func CreateUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	// cuando usamos new para decodificar la entrada json estamos usando el modelo completo incluyendo la pass, con lo cual tenemos que cifrarla
	user := new(models.User)

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	// comprobar que no pase campos vacios
	if user.Name == "" || user.Email == "" || user.Username == "" || user.Password == "" {
		http.Error(w, "all fields are required", http.StatusBadRequest)
		return
	}

	// ciframos la pass que creo el usuario
	password, err := hashPw(user.Password)
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	// le asignamos la pass cifrada
	user.Password = password

	if err := db.Create(&user).Error; err != nil {
		http.Error(w, "error creating the user", http.StatusInternalServerError)
		return
	}

	// ponemos la pass a vacio para mostrarle al usuario despues de que haya sido creado
	user.Password = ""

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// funcion para cifrar la pass al momento de crear un user
func hashPw(password string) (string, error) {
	pass, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(pass), err
}
