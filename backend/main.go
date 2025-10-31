package main

import (
	"backend/api"
	"backend/db"
	"fmt"
)

func main() {
	connection := db.Connect()
	defer connection.Close()

	db.RunMigrations(connection)

	if err := api.ImportCursos(connection); err != nil {
		fmt.Println("❌ Error importando datos:", err)
	}
}
