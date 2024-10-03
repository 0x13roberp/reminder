package middleware

import (
	"encoding/json"
	"goapi/config"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte(config.JwtKey)

// si surge algun error durante la validacion. estilo http.Error
func jwtErrorMessage(w http.ResponseWriter, err string, status int) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{"error": err})
}

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			jwtErrorMessage(w, "missing or malformed jwt", http.StatusBadRequest)
			return
		}

		// quitamos el Bearer del header para quedarnos con el token.
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			jwtErrorMessage(w, "missing or malformed jwt", http.StatusBadRequest)
			return
		}

		// extraer la informacion que tiene el token.
		claims := &jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, http.ErrNotSupported
			}
			return []byte(config.JwtKey), nil
		})

		if err != nil {
			jwtErrorMessage(w, "invalid jwt", http.StatusUnauthorized)
			return
		}

		if !token.Valid {
			jwtErrorMessage(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		// si el token es valido y la autenticacion fue exitosa el middleware llama al siguiente handler.
		next.ServeHTTP(w, r)
	})
}

func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("you can access to the route"))
}
