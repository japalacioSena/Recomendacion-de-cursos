package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

// Connect abre la conexión a PostgreSQL y la retorna
func Connect() *sql.DB {
	// Si no hay variable de entorno, usa db:5432 (Docker)
	dbHost := getEnv("DB_HOST", "db")
	dbPort := getEnv("DB_PORT", "5432") // dentro del contenedor es siempre 5432
	dbUser := getEnv("DB_USER", "recomendaciones")
	dbPassword := getEnv("DB_PASSWORD", "recomendaciones_postgres_589")
	dbName := getEnv("DB_NAME", "recomendaciones_user_postgres")

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("❌ Error creating DB connection: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("🚫 Cannot connect to PostgreSQL: %v", err)
	}

	fmt.Println("✅ Connected to PostgreSQL successfully!")
	return db
}

// getEnv obtiene una variable de entorno con valor por defecto
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
