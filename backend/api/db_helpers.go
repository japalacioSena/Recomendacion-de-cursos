package api

import (
	"context"
	"database/sql"
	"fmt"
)

// GetOrCreateID obtiene el ID de un registro existente o lo crea si no existe.
// Retorna un error si el valor es vacío o si ocurre un problema con la base de datos.
func GetOrCreateID(ctx context.Context, db *sql.DB, table string, column string, value interface{}) (int, error) {
	// Validar valores vacíos o nulos
	if value == nil {
		return 0, fmt.Errorf("⚠️ valor nulo para %s.%s", table, column)
	}

	switch v := value.(type) {
	case string:
		if v == "" {
			return 0, fmt.Errorf("⚠️ valor vacío para %s.%s", table, column)
		}
	}

	// Intentar obtener el ID si ya existe
	var id int
	querySelect := fmt.Sprintf("SELECT id FROM %s WHERE %s = $1", table, column)
	err := db.QueryRowContext(ctx, querySelect, value).Scan(&id)

	if err != nil && err != sql.ErrNoRows {
		return 0, fmt.Errorf("error buscando en %s.%s (%v): %w", table, column, value, err)
	}

	// Si existe, devolver el ID encontrado
	if err == nil {
		return id, nil
	}

	// Si no existe, insertar el nuevo registro y devolver su ID
	queryInsert := fmt.Sprintf("INSERT INTO %s (%s) VALUES ($1) RETURNING id", table, column)
	err = db.QueryRowContext(ctx, queryInsert, value).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("error insertando en %s (%s=%v): %w", table, column, value, err)
	}

	return id, nil
}
