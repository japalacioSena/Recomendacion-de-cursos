package importzajuna

import (
	"context"
	"database/sql"
	"fmt"
)

func ImportUsers(localDB *sql.DB) error {

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

	return nil
}
