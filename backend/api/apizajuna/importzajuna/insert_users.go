package importzajuna

import (
	"context"
	"database/sql"
	"fmt"
)

// insertUsuarioZajuna inserta un usuarios con sus relaciones asegurando las foreign keys.
func insertUserZajuna(ctx context.Context, db *sql.DB, u UserZajuna) error {
	fmt.Printf("🧩 PrfDuracionMaxima = %v \n", db)

	// Insertar los datos de los usuarios
	_, err := db.ExecContext(ctx, `
	INSERT INTO user_zajuna (username, idnumber, firstname, lastname)
	VALUES ($1, $2, $3, $4)
	ON CONFLICT (username) DO UPDATE 
	SET idnumber = EXCLUDED.idnumber,
		firstname = EXCLUDED.firstname,
		lastname = EXCLUDED.lastname
`, u.Username, u.Idnumber, u.Firstname, u.Lastname)

	if err != nil {
		return fmt.Errorf("❌ error insertando usuario %v: %v", u.Username, err)
	}

	fmt.Printf("✅ Usuario %v insertado/actualizado correctamente\n", u.Username)
	return nil

}
