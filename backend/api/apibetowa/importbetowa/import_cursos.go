package importbetowa

import (
	"context"
	"database/sql"
	"fmt"
)

func ImportCursos(localDB *sql.DB) error {
	fmt.Println("📥 Importando cursos desde base externa...")

	cursos, err := GetCursoBetowa()
	if err != nil {
		return fmt.Errorf("error al obtener cursos externos: %v", err)
	}

	ctx := context.Background()

	for _, c := range cursos {
		err := insertCursoBetowa(ctx, localDB, c)
		if err != nil {
			fmt.Println("❌", err)
			continue
		}
	}

	return nil
}
