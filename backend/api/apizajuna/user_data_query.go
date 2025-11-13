package apizajuna

import (
	"backend/db"
	"encoding/json"
	"log"
	"net/http"
)

// Estructura de respuesta
type UserResponse struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Handler para consultar usuario “quemado”
func GetFixedUserHandler(w http.ResponseWriter, r *http.Request) {
	connection := db.Connect()
	defer connection.Close()

	username := "29873380" // 🔥 Usuario quemado

	var user UserResponse
	query := `SELECT firstname, lastname FROM user_zajuna WHERE username = $1 LIMIT 1`

	err := connection.QueryRow(query, username).Scan(&user.Firstname, &user.Lastname)
	if err != nil {
		http.Error(w, "Usuario no encontrado o error en la consulta", http.StatusNotFound)
		log.Println("❌ Error consultando usuario:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
