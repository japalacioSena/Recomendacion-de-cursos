package main

import (
	"backend/api/apibetowa/importbetowa"
	"backend/api/apizajuna/importzajuna"
	"backend/api/query"
	"backend/db"
	"backend/db/betowa"
	"backend/db/zajuna"
	"backend/utils"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	connection := db.Connect()
	defer connection.Close()

	betowa.RunMigrations_betowa(connection)
	zajuna.RunMigrations_zajuna(connection)

	// Rutas del API
	http.HandleFunc("/api/user_zajuna", utils.EnableCORS(query.GetFixedUserHandler))

	fmt.Println("🚀 Servidor corriendo en puerto 8080")
	go func() {
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()

	if err := importbetowa.ImportCursos(connection); err != nil {
		fmt.Println("❌ Error importando datos:", err)
	}

	// Tareas programadas
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for {
		// Importar desde Betowa
		log.Println("📥 Ejecutando importación desde Betowa...")
		connection := db.Connect()
		if err := importbetowa.ImportCursos(connection); err != nil {
			log.Println("❌ Error importando datos:", err)
		} else {
			log.Println("✅ Importación completada Betowa")
		}
		connection.Close()

		// Importar desde Zajuna
		log.Println("📥 Ejecutando importación desde Zajuna...")
		connection = db.Connect()
		if err := importzajuna.ImportUsers(connection); err != nil {
			log.Println("❌ Error importando datos Zajuna:", err)
		} else {
			log.Println("✅ Importación completada Zajuna")
		}
		connection.Close()

		// Espera hasta la próxima iteración
		<-ticker.C
	}
}
