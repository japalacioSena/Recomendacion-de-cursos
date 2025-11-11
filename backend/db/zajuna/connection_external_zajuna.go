package zajuna

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// ConnectExternal establece una conexión a una base de datos PostgreSQL de betowa
func ConnectExternalZajuna() *sql.DB {
	host := "192.168.1.122"
	port := 5432
	user := "postgres"
	password := "12345"
	dbname := "zajuna"

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("❌ Error abriendo conexión a zajuna: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("🚫 No se pudo conectar a zajuna: %v", err)
	}

	log.Println("✅ Conectado correctamente a la base externa zajuna")
	return db
}
