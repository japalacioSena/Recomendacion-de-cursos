package apizajuna

import (
	"context"
	"database/sql"
	"fmt"
)

func ImportCursos(localDB *sql.DB) error {
	fmt.Println("📥 Importando usuarios desde base externa...")

	userZajuna, err := GetUserZajuna()
	if err != nil {
		return fmt.Errorf("error al obtener usuarios de zajuna: %v", err)
	}

	ctx := context.Background()

	for _, u := range userZajuna {
		err := insertUserZajuna(ctx, localDB, u)
		if err != nil {
			fmt.Println("❌", err)
			continue
		}
	}

	fmt.Printf("✅ Se importaron %d usuarios correctamente\n", len(userZajuna))
	return nil
}
