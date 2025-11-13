package importzajuna

import (
	"backend/db/zajuna"
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
)

// Estructura de los datos que vienen desde Betowa
type UserZajuna struct {
	Username  string
	Idnumber  string
	Firstname string
	Lastname  string
}

// GetUserZajuna obtiene los UserZajuna desde la base externa Betowa
func GetUserZajuna() ([]UserZajuna, error) {
	connection := GetUserZajunaDBConnection()
	if connection == nil {
		return nil, fmt.Errorf("no se pudo conectar a la base de datos externa")
	}
	defer connection.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `SELECT username,idnumber, firstname,lastname FROM mdl_user`

	rows, err := connection.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("error ejecutando la consulta: %v", err)
	}
	defer rows.Close()

	var userZajuna []UserZajuna
	for rows.Next() {
		var u UserZajuna
		err := rows.Scan(
			&u.Username, &u.Idnumber, &u.Firstname, &u.Lastname,
		)
		if err != nil {
			return nil, fmt.Errorf("error escaneando fila: %v", err)
		}
		userZajuna = append(userZajuna, u)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error en las filas: %v", err)
	}

	log.Printf("✅ Se obtuvieron %d usuarios desde zajuna\n", len(userZajuna))
	return userZajuna, nil
}

// Conexión a la base de datos Betowa
func GetUserZajunaDBConnection() *sql.DB {
	return zajuna.ConnectExternalZajuna()
}
