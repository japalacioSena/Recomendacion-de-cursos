package zajuna

import (
	"database/sql"
	"log"
)

func RunMigrations_zajuna(db *sql.DB) {
	tables := []string{
		`CREATE TABLE IF NOT EXISTS user_zajuna (
			id SERIAL PRIMARY KEY,
			username TEXT UNIQUE,
			idnumber tEXT UNIQUE, 
			firstname TEXT,
			lastname TEXT
		)`,
		`CREATE TABLE IF NOT EXISTS course_tracking (
			id SERIAL PRIMARY KEY,
			id_user INT,
			id_technological_red INT,
			FOREIGN KEY (id_user) REFERENCES user_zajuna(id),
			FOREIGN KEY (id_technological_red) REFERENCES technological_red(id)
		)`,
	}

	for _, table := range tables {
		if _, err := db.Exec(table); err != nil {
			log.Fatalf("❌ Error creando tabla: %v", err)
		}
	}

	log.Println("✅ Migraciones Zajuna ejecutadas correctamente")
}
