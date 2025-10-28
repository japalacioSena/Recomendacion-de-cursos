package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	// Obtener variables de entorno
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "5432")
	dbUser := getEnv("DB_USER", "postgres")
	dbPassword := getEnv("DB_PASSWORD", "postgres")
	dbName := getEnv("DB_NAME", "postgres")

	// Mostrar la configuración (sin mostrar la contraseña completa por seguridad)
	fmt.Println("🔧 Database configuration:")
	fmt.Printf("  Host: %s\n", dbHost)
	fmt.Printf("  Port: %s\n", dbPort)
	fmt.Printf("  User: %s\n", dbUser)
	fmt.Printf("  DB:   %s\n", dbName)

	// Construir el string de conexión
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName,
	)

	// Conectar a PostgreSQL
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("❌ Error creating DB connection: %v", err)
	}
	defer db.Close()

	// Verificar conexión
	err = db.Ping()
	if err != nil {
		log.Fatalf("🚫 Cannot connect to PostgreSQL: %v", err)
	}

	fmt.Println("✅ Connected to PostgreSQL successfully!")
}

// getEnv obtiene una variable de entorno con valor por defecto
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
