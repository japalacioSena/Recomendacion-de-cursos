package importbetowa

import (
	"backend/db/betowa"
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
)

// Estructura de los datos que vienen desde Betowa
type Curso struct {
	PrfCodigo            int
	TipoDeFormacion      string
	PrfDenominacion      string
	NivelDeFormacion     string
	PrfDuracionMaxima    int
	PrfDurEtapaLectiva   int
	PrfDurEtapaProd      int
	PrfFchRegistro       string
	FechaActivo          string
	PrfDescripcionReq    string
	PrfCreditos          int
	LineaTecnologica     string
	RedTecnologica       string
	RedConocimiento      string
	Modalidad            string
	ApuestasPrioritarias string
	Indice               string
	Ocupacion            string
}

// GetCursoBetowa obtiene los cursos desde la base externa Betowa
func GetCursoBetowa() ([]Curso, error) {
	connection := GetCursoBetowaDBConnection()
	if connection == nil {
		return nil, fmt.Errorf("no se pudo conectar a la base de datos externa")
	}
	defer connection.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `SELECT PRF_CODIGO, TIPO_DE_FORMACION, PRF_DENOMINACION, NIVEL_DE_FORMACION,
		PRF_DURACION_MAXIMA, PRF_DUR_ETAPA_LECTIVA, PRF_DUR_ETAPA_PROD, PRF_FCH_REGISTRO, Fecha_Activo,
		PRF_DESCRIPCION_REQUISITO, PRF_CREDITOS, Linea_Tecnologica, Red_Tecnologica, Red_Conocimiento,
		Modalidad, apuestas_prioritarias, Indice, Ocupacion FROM cursos`

	rows, err := connection.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("error ejecutando la consulta: %v", err)
	}
	defer rows.Close()

	var cursos []Curso
	for rows.Next() {
		var c Curso
		err := rows.Scan(
			&c.PrfCodigo, &c.TipoDeFormacion, &c.PrfDenominacion, &c.NivelDeFormacion,
			&c.PrfDuracionMaxima, &c.PrfDurEtapaLectiva, &c.PrfDurEtapaProd, &c.PrfFchRegistro,
			&c.FechaActivo, &c.PrfDescripcionReq, &c.PrfCreditos, &c.LineaTecnologica,
			&c.RedTecnologica, &c.RedConocimiento, &c.Modalidad, &c.ApuestasPrioritarias,
			&c.Indice, &c.Ocupacion,
		)
		if err != nil {
			return nil, fmt.Errorf("error escaneando fila: %v", err)
		}
		cursos = append(cursos, c)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error en las filas: %v", err)
	}

	log.Printf("✅ Se obtuvieron %d cursos desde Betowa\n", len(cursos))
	return cursos, nil
}

// Conexión a la base de datos Betowa
func GetCursoBetowaDBConnection() *sql.DB {
	return betowa.ConnectExternal()
}
