package apibetowa

import (
	"backend/api"
	"context"
	"database/sql"
	"fmt"
)

// insertCursoBetowa inserta un curso con sus relaciones asegurando las foreign keys.
func insertCursoBetowa(ctx context.Context, db *sql.DB, c Curso) error {
	fmt.Printf("🧩 PrfDuracionMaxima = %v \n", db)
	// Helper inline para reducir repetición
	get := func(table, column string, value interface{}) (int, error) {
		id, err := api.GetOrCreateID(ctx, db, table, column, value)
		if err != nil {
			return 0, fmt.Errorf("error en %s.%s (valor=%v) para curso %v: %v", table, column, value, c.PrfCodigo, err)
		}
		return id, nil
	}

	// Obtener o crear los IDs relacionados
	idType, err := get("type_of_training", "name", c.TipoDeFormacion)
	if err != nil {
		return err
	}

	idDenomination, err := get("denomination", "name", c.PrfDenominacion)
	if err != nil {
		return err
	}

	idLevel, err := get("level_of_training", "name", c.NivelDeFormacion)
	if err != nil {
		return err
	}

	fmt.Printf("🧩 PrfDuracionMaxima = %v (tipo=%T)\n", c.PrfDuracionMaxima, c.PrfDuracionMaxima)
	idMaxDuration, err := get("maximum_duration", "duration", c.PrfDuracionMaxima)
	if err != nil {
		return err
	}

	idLective, err := get("duration_academic_stage", "duration", c.PrfDurEtapaLectiva)
	if err != nil {
		return err
	}

	idProductive, err := get("duration_production_stage", "duration", c.PrfDurEtapaProd)
	if err != nil {
		return err
	}

	idRegDate, err := get("registration_date", "date", c.PrfFchRegistro)
	if err != nil {
		return err
	}

	idActiveDate, err := get("active_date", "date", c.FechaActivo)
	if err != nil {
		return err
	}

	idReqDesc, err := get("description_requirements", "description", c.PrfDescripcionReq)
	if err != nil {
		return err
	}

	idCredits, err := get("credits", "credits", c.PrfCreditos)
	if err != nil {
		return err
	}

	idTechLine, err := get("technological_line", "name", c.LineaTecnologica)
	if err != nil {
		return err
	}

	idTechRed, err := get("technological_red", "name", c.RedTecnologica)
	if err != nil {
		return err
	}

	idKnowNet, err := get("knowledge_network", "name", c.RedConocimiento)
	if err != nil {
		return err
	}

	idMode, err := get("mode", "name", c.Modalidad)
	if err != nil {
		return err
	}

	idPriority, err := get("priority_bets", "name", c.ApuestasPrioritarias)
	if err != nil {
		return err
	}

	idIndex, err := get("index_cursos", "name", c.Indice)
	if err != nil {
		return err
	}

	idOccupation, err := get("occupational_profile", "description", c.Ocupacion)
	if err != nil {
		return err
	}

	// Insertar el curso con las foreign keys
	_, err = db.ExecContext(ctx, `
		INSERT INTO cursos (
			course_code,
			id_type_of_training,
			id_denomination,
			id_level_of_training,
			id_maximum_duration,
			id_duration_academic_stage,
			id_duration_production_stage,
			id_registration_date,
			id_active_date,
			id_description_requirements,
			id_credits,
			id_technological_line,
			id_technological_red,
			id_knowledge_network,
			id_mode,
			id_priority_bets,
			id_index_cursos,
			id_occupational_profile
		) VALUES (
			$1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18
		)
		ON CONFLICT (course_code) DO NOTHING
	`,
		c.PrfCodigo,
		idType, idDenomination, idLevel, idMaxDuration, idLective, idProductive,
		idRegDate, idActiveDate, idReqDesc, idCredits, idTechLine, idTechRed,
		idKnowNet, idMode, idPriority, idIndex, idOccupation,
	)

	if err != nil {
		return fmt.Errorf("❌ error insertando curso %v: %v", c.PrfCodigo, err)
	}

	fmt.Printf("✅ Curso %v insertado correctamente\n", c.PrfCodigo)
	return nil
}
