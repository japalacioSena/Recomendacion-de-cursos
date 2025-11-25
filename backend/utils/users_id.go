// En el archivo: src/utils/http.go
package utils

import (
	"errors"
	"net/http"
	"strconv"
)

// Definición de errores. Solo devuelve el error sin enviar la respuesta HTTP.
var ErrMissingUserParameter = errors.New("falta el parámetro id_user")
var ErrInvalidUserParameter = errors.New("el parámetro id_user debe ser un número entero válido")

// GetUserIDFromRequest extrae y valida el parámetro "id_user" de la URL.
// NOTA: Esta versión NO recibe 'http.ResponseWriter'.
func GetUserIDFromRequest(r *http.Request) (int, error) {
	// 1. Obtener el parámetro
	idUserStr := r.URL.Query().Get("id_user")

	// 2. Validar presencia
	if idUserStr == "" {
		// Devuelve nuestro error personalizado
		return 0, ErrMissingUserParameter
	}

	// 3. Convertir a entero
	idUser, err := strconv.Atoi(idUserStr)
	if err != nil {
		// Devuelve un error si la conversión falla
		return 0, ErrInvalidUserParameter
	}

	return idUser, nil
}
