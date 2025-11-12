package main

import (
	"backend/api/apibetowa"
	"backend/api/apizajuna"
	"backend/db"
	"backend/db/betowa"
	"backend/db/zajuna"
	"fmt"
	"log"
	"time"
)

func main() {
	connection := db.Connect()
	defer connection.Close()

	betowa.RunMigrations_betowa(connection)
	zajuna.RunMigrations_zajuna(connection)

	if err := apibetowa.ImportCursos(connection); err != nil {
		fmt.Println("❌ Error importando datos:", err)
	}

	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for {
		// Importar desde Betowa
		log.Println("📥 Ejecutando importación desde Betowa...")
		connection := db.Connect()
		if err := apibetowa.ImportCursos(connection); err != nil {
			log.Println("❌ Error importando datos:", err)
		} else {
			log.Println("✅ Importación completada Betowa")
		}
		connection.Close()

		// Importar desde Zajuna
		log.Println("📥 Ejecutando importación desde Zajuna...")
		connection = db.Connect()
		if err := apizajuna.ImportUsers(connection); err != nil {
			log.Println("❌ Error importando datos Zajuna:", err)
		} else {
			log.Println("✅ Importación completada Zajuna")
		}
		connection.Close()

		// Espera hasta la próxima iteración
		<-ticker.C
	}

}
