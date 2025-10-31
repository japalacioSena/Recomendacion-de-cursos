package api

import (
	"context"
	"database/sql"
	"fmt"
)

// ImportCursos importa los cursos desde la base externa hacia la base local
func ImportCursos(localDB *sql.DB) error {
	fmt.Println("📥 Importando cursos desde base externa...")

	cursos, err := GetCursoBetowa()
	if err != nil {
		return fmt.Errorf("error al obtener cursos externos: %v", err)
	}

	ctx := context.Background()

	for _, c := range cursos {
		_, err := localDB.ExecContext(ctx, `
			INSERT INTO cursos (
			course_code, type_of_training, denomination, training_level,
			max_duration, lective_duration, productive_duration, registration_date,
			active_date, requirement_desc, credits, tech_line,
			tech_network, knowledge_network, modality, priority_bets,
			index, occupation
		) VALUES (
			$1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18
		)
		ON CONFLICT (course_code) DO NOTHING
	`,
			c.PrfCodigo, c.TipoDeFormacion, c.PrfDenominacion, c.NivelDeFormacion,
			c.PrfDuracionMaxima, c.PrfDurEtapaLectiva, c.PrfDurEtapaProd, c.PrfFchRegistro,
			c.FechaActivo, c.PrfDescripcionReq, c.PrfCreditos, c.LineaTecnologica,
			c.RedTecnologica, c.RedConocimiento, c.Modalidad, c.ApuestasPrioritarias,
			c.Indice, c.Ocupacion,
		)
		if err != nil {
			fmt.Println("❌ Error insertando curso:", err)
			continue
		}
	}

	fmt.Printf("✅ Se importaron %d cursos correctamente\n", len(cursos))
	return nil
}
