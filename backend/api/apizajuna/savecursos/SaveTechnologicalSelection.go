package savecursos

import (
	"backend/db"
	"encoding/json"
	"fmt"
	"net/http"
)

func SaveTechnologicalSelectionHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		UserID int   `json:"user_id"`
		RedIDs []int `json:"red_ids"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Error en el JSON recibido", http.StatusBadRequest)
		return
	}

	conn := db.Connect()
	defer conn.Close()

	// ---------------------------------------------------------------------
	// 1️⃣ Eliminar lo que ya tenga registrado el usuario
	// ---------------------------------------------------------------------
	_, err := conn.Exec(`DELETE FROM course_tracking WHERE id_user = $1`, req.UserID)
	if err != nil {
		http.Error(w, "Error limpiando datos anteriores", http.StatusInternalServerError)
		fmt.Println("❌ Error DELETE:", err)
		return
	}

	// ---------------------------------------------------------------------
	// 2️⃣ Insertar las nuevas selecciones
	// ---------------------------------------------------------------------
	for _, redID := range req.RedIDs {
		_, err := conn.Exec(
			`INSERT INTO course_tracking (id_user, id_technological_red) VALUES ($1, $2)`,
			req.UserID, redID,
		)
		if err != nil {
			http.Error(w, "Error guardando selección", http.StatusInternalServerError)
			fmt.Println("❌ Error INSERT:", err)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Selección actualizada correctamente"))
}
