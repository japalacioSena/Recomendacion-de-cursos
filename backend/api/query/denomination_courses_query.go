package query

import (
	"backend/db"
	"backend/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

// Estructura de respuesta (alineada con las 4 columnas seleccionadas)
type DenominationCursosResponse struct {
	ID               int    `json:"id"`                // co.course_code (convertido)
	Name             string `json:"name"`              // de.name
	RegistrationDate string `json:"registration_date"` // rd.date (formateado)
	ActiveDate       string `json:"active_date"`       // ad.date (formateado)
}

// Handler para consultar los cursos
func GetDenominationCursosHandler(w http.ResponseWriter, r *http.Request) {
	connection := db.Connect()
	defer connection.Close()

	idUser, err := utils.GetUserIDFromRequest(r)

	// ... (Manejo de error idUser) ...
	if err != nil {
		statusCode := http.StatusBadRequest
		// Puedes simplificar esto ya que utils.ErrMissingUserParameter y utils.ErrInvalidUserParameter
		// son los únicos errores esperados
		http.Error(w, err.Error(), statusCode)
		return
	}

	// Consulta: Se eliminó GROUP BY y se agregó un CAST para simplificar la vida.
	query := `
        select c.course_code, d.name, rd.date, ad.date from denomination d join cursos c on c.id_denomination = d.id
	join registration_date rd on c.id_registration_date = rd.id 
	join active_date ad on c.id_active_date = ad.id 
	where c.id_technological_red IN(
	select ct.id_technological_red  from course_tracking ct
	where ct.id_user = $1)
        ORDER BY de.name
    `

	rows, err := connection.Query(query, idUser)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error ejecutando la consulta: %v", err), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var cursos []DenominationCursosResponse

	for rows.Next() {
		var c DenominationCursosResponse

		// Variables para escanear las columnas de la BD
		var courseCodeStr string // course_code es probablemente VARCHAR/TEXT
		var name string
		var regDate, actDate time.Time // Las fechas se reciben como time.Time

		// 🚨 ESCANEO CORRECTO: 4 COLUMNAS A 4 VARIABLES 🚨
		err := rows.Scan(&courseCodeStr, &name, &regDate, &actDate)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error leyendo datos: %v", err), http.StatusInternalServerError)
			return
		}

		// --- Mapeo y Conversión ---

		// 1. Convertir course_code (string) a ID (int)
		c.ID, err = strconv.Atoi(courseCodeStr)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error de conversión de ID: %v", err), http.StatusInternalServerError)
			return
		}

		c.Name = name
		// 2. Formatear las fechas para JSON (usando el formato ISO)
		c.RegistrationDate = regDate.Format("2006-01-02")
		c.ActiveDate = actDate.Format("2006-01-02")

		cursos = append(cursos, c)
	}

	// Verificación final de errores del iterador
	if err = rows.Err(); err != nil {
		http.Error(w, fmt.Sprintf("Error durante la iteración: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cursos)
}
