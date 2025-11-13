package query

import (
	"backend/db"
	"encoding/json"
	"log"
	"net/http"
)

// Estructura de respuesta (exportada con mayúsculas)
type TechnologicalRedResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Handler para consultar la red tecnológica
func GetTechnologicalRedHandler(w http.ResponseWriter, r *http.Request) {
	// Conectar a la base de datos
	connection := db.Connect()
	defer connection.Close()

	// Ejecutar la consulta
	query := `SELECT id, name FROM technological_red ORDER BY name`

	rows, err := connection.Query(query)
	if err != nil {
		http.Error(w, "Error ejecutando la consulta", http.StatusInternalServerError)
		log.Println("❌ Error ejecutando la consulta:", err)
		return
	}
	defer rows.Close()

	// Crear slice para almacenar los resultados
	var redes []TechnologicalRedResponse

	for rows.Next() {
		var red TechnologicalRedResponse
		if err := rows.Scan(&red.ID, &red.Name); err != nil {
			http.Error(w, "Error leyendo los datos", http.StatusInternalServerError)
			log.Println("❌ Error leyendo fila:", err)
			return
		}
		redes = append(redes, red)
	}

	// Verificar si no se obtuvieron resultados
	if len(redes) == 0 {
		http.Error(w, "No se encontraron redes tecnológicas", http.StatusNotFound)
		return
	}

	// Devolver respuesta en formato JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(redes); err != nil {
		http.Error(w, "Error generando la respuesta JSON", http.StatusInternalServerError)
		log.Println("❌ Error codificando JSON:", err)
	}
}
