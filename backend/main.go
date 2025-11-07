package main

import (
	"backend/api/apibetowa"
	"backend/db"
	"backend/db/betowa"
	"fmt"
	"log"
	"time"
)

func main() {
	connection := db.Connect()
	defer connection.Close()

	betowa.RunMigrations_betowa(connection)

	if err := apibetowa.ImportCursos(connection); err != nil {
		fmt.Println("❌ Error importando datos:", err)
	}

	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for {
		log.Println("📥 Ejecutando importación desde Betowa...")
		connection := db.Connect()
		if err := apibetowa.ImportCursos(connection); err != nil {
			log.Println("❌ Error importando datos:", err)
		} else {
			log.Println("✅ Importación completada")
		}
		connection.Close()

		// Espera hasta la próxima iteración
		<-ticker.C
	}
}
