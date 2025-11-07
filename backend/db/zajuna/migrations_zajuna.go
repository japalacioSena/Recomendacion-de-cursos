package zajuna

import (
	"database/sql"
	"fmt"
	"log"
)

func RunMigrations_betowa(db *sql.DB) {
	tables := []string{
		`CREATE TABLE IF NOT EXISTS user (
			id SERIAL PRIMARY KEY,
			code INT UNIQUE
		)`,
		`CREATE TABLE IF NOT EXISTS mailing _list_course (
			id SERIAL PRIMARY KEY,
			name TEXT UNIQUE
		)`,
		`CREATE TABLE IF NOT EXISTS course_users (
			id SERIAL PRIMARY KEY,
			id_course INT,
			id_user INT,
			foreign key (id_course) references mailing_list_course(id),
			foreign key (id_user) references user(id)
		)`,
	}

	for _, query := range tables {
		_, err := db.Exec(query)
		if err != nil {
			log.Fatalf("❌ Error en migración:\n%v\nQuery: %s\n", err, query)
		}
	}

	fmt.Println("✅ Migraciones completadas correctamente.")

}
