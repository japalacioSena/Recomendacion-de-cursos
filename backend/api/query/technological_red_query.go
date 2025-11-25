package query

import (
	"backend/db"
	"backend/utils"
	"encoding/json"
	"errors"
	"net/http"
)

// Estructura de respuesta (exportada con mayúsculas)
type TechnologicalRedResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Selected bool   `json:"selected"`
}

// Handler para consultar la red tecnológica
func GetTechnologicalRedHandler(w http.ResponseWriter, r *http.Request) {
	connection := db.Connect()
	defer connection.Close()

	idUser, err := utils.GetUserIDFromRequest(r)

	if err != nil {
		// Asume un estado 400 (Bad Request) para errores de validación de entrada
		statusCode := http.StatusBadRequest

		// Opcional: Podrías diferenciar errores, pero 400 es adecuado para ambos casos.
		if errors.Is(err, utils.ErrMissingUserParameter) {
			// Este es un ejemplo para ser explícito, pero usa el mensaje de error.
		} else if errors.Is(err, utils.ErrInvalidUserParameter) {
			// ...
		}

		http.Error(w, err.Error(), statusCode)
		return // ¡Detiene la ejecución si hay un error!
	}

	query := `
		SELECT tr.id, tr.name,
		CASE WHEN ct.id_technological_red IS NULL THEN false ELSE true END AS selected
		FROM technological_red tr
		LEFT JOIN course_tracking ct 
			ON tr.id = ct.id_technological_red 
			AND ct.id_user = $1
		ORDER BY tr.name
	`

	rows, err := connection.Query(query, idUser)

	if err != nil {
		http.Error(w, "Error ejecutando la consulta", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	type Response struct {
		ID       int    `json:"id"`
		Name     string `json:"name"`
		Selected bool   `json:"selected"`
	}

	var redes []Response

	for rows.Next() {
		var r Response
		err := rows.Scan(&r.ID, &r.Name, &r.Selected)
		if err != nil {
			http.Error(w, "Error leyendo datos", http.StatusInternalServerError)
			return
		}
		redes = append(redes, r)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(redes)
}
