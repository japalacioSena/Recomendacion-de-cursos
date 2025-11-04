package db

import (
	"database/sql"
	"fmt"
	"log"
)

func RunMigrations(db *sql.DB) {
	tables := []string{
		`CREATE TABLE IF NOT EXISTS type_of_training (
			id SERIAL PRIMARY KEY,
			name TEXT UNIQUE
		)`,
		`CREATE TABLE IF NOT EXISTS denomination (
			id SERIAL PRIMARY KEY,
			name TEXT UNIQUE
		)`,
		`CREATE TABLE IF NOT EXISTS level_of_training (
			id SERIAL PRIMARY KEY,
			name TEXT UNIQUE
		)`,
		`CREATE TABLE IF NOT EXISTS maximum_duration (
			id SERIAL PRIMARY KEY,
			duration INT UNIQUE
		)`,
		`CREATE TABLE IF NOT EXISTS duration_academic_stage (
			id SERIAL PRIMARY KEY,
			duration INT UNIQUE
		)`,
		`CREATE TABLE IF NOT EXISTS duration_production_stage (
			id SERIAL PRIMARY KEY,
			duration INT UNIQUE
		)`,
		`CREATE TABLE IF NOT EXISTS registration_date (
			id SERIAL PRIMARY KEY,
			date DATE UNIQUE
		)`,
		`CREATE TABLE IF NOT EXISTS active_date (
			id SERIAL PRIMARY KEY,
			date DATE UNIQUE
		)`,
		`CREATE TABLE IF NOT EXISTS description_requirements (
			id SERIAL PRIMARY KEY,
			description TEXT UNIQUE
		)`,
		`CREATE TABLE IF NOT EXISTS credits (
			id SERIAL PRIMARY KEY,
			credits INT UNIQUE
		)`,
		`CREATE TABLE IF NOT EXISTS technological_line (
			id SERIAL PRIMARY KEY,
			name TEXT UNIQUE
		)`,
		`CREATE TABLE IF NOT EXISTS technological_red (
			id SERIAL PRIMARY KEY,
			name TEXT UNIQUE
		)`,
		`CREATE TABLE IF NOT EXISTS knowledge_network (
			id SERIAL PRIMARY KEY,
			name TEXT UNIQUE
		)`,
		`CREATE TABLE IF NOT EXISTS mode (
			id SERIAL PRIMARY KEY,
			name TEXT UNIQUE
		)`,
		`CREATE TABLE IF NOT EXISTS priority_bets (
			id SERIAL PRIMARY KEY,
			name TEXT UNIQUE
		)`,
		`CREATE TABLE IF NOT EXISTS index_cursos (
			id SERIAL PRIMARY KEY,
			name TEXT UNIQUE
		)`,
		`CREATE TABLE IF NOT EXISTS occupational_profile (
			id SERIAL PRIMARY KEY,
			description TEXT UNIQUE
		)`,
		`CREATE TABLE IF NOT EXISTS cursos (
			id SERIAL PRIMARY KEY,
			course_code INT UNIQUE,
			id_type_of_training INT,
			id_denomination INT,
			id_level_of_training INT,
			id_maximum_duration INT,
			id_duration_academic_stage INT,
			id_duration_production_stage INT,
			id_registration_date INT,
			id_active_date INT,
			id_description_requirements INT,
			id_credits INT,
			id_technological_line INT,
			idtechnological_red INT,
			id_knowledge_network INT,
			id_mode INT,
			id_priority_bets INT,
			id_index_cursos INT,
			id_occupational_profile INT,
			FOREIGN KEY (id_type_of_training) REFERENCES type_of_training(id),
			FOREIGN KEY (id_denomination) REFERENCES denomination(id),
			FOREIGN KEY (id_level_of_training) REFERENCES level_of_training(id),
			FOREIGN KEY (id_maximum_duration) REFERENCES maximum_duration(id),
			FOREIGN KEY (id_duration_academic_stage) REFERENCES duration_academic_stage(id),
			FOREIGN KEY (id_duration_production_stage) REFERENCES duration_production_stage(id),
			FOREIGN KEY (id_registration_date) REFERENCES registration_date(id),
			FOREIGN KEY (id_active_date) REFERENCES active_date(id),
			FOREIGN KEY (id_description_requirements) REFERENCES description_requirements(id),
			FOREIGN KEY (id_credits) REFERENCES credits(id),
			FOREIGN KEY (id_technological_line) REFERENCES technological_line(id),
			FOREIGN KEY (idtechnological_red) REFERENCES technological_red(id),
			FOREIGN KEY (id_knowledge_network) REFERENCES knowledge_network(id),
			FOREIGN KEY (id_mode) REFERENCES mode(id),
			FOREIGN KEY (id_priority_bets) REFERENCES priority_bets(id),
			FOREIGN KEY (id_index_cursos) REFERENCES index_cursos(id),
			FOREIGN KEY (id_occupational_profile) REFERENCES occupational_profile(id)
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
