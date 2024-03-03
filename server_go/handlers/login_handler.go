package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Credentials struct {
	Usuario string `json:"usuario"`
	Clave   string `json:"clave"`
}

func Login_handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Metodo no permitido", http.StatusMethodNotAllowed)
		return
	}

	var credencial Credentials
	err := json.NewDecoder(r.Body).Decode(&credencial)
	if err != nil {
		http.Error(w, "Error: No se convertido la solicitud HTTP a formato JSON", http.StatusBadRequest)
		return
	}

	if credencial.Usuario == "" || credencial.Clave == "" {
		http.Error(w, "Los atributos 'usuario' y 'clave' son obligatorios", http.StatusBadRequest)
		return
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = credencial.Usuario
	claims["exp"] = time.Now().Add(time.Hour).Unix()
	claims["iss"] = "ingesis.uniquindio.edu.co"

	tokenString, err := token.SignedString([]byte("12345"))
	if err != nil {
		http.Error(w, "Error al firmar el token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprint(w, tokenString)
}
