// Obtiene la información del usuario
export async function getUser() {
  const res = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/api/user_zajuna`);
  if (!res.ok) throw new Error("Error obteniendo usuario");
  return res.json();
}

// Obtiene las redes tecnológicas disponibles
export async function getTechnologicalRed(idUser) {
  const res = await fetch(
    `${process.env.NEXT_PUBLIC_API_URL}/api/technological_reds?id_user=${idUser}`
  );

  if (!res.ok) throw new Error("Error obteniendo redes tecnológicas");
  return res.json();
}


// Guarda la selección tecnológica del usuario
export async function SaveTechnologicalSelectionHandler(userId, redIds) {
  console.log("Guardando selección tecnológica para el usuario:", userId, "con redes:", redIds);
  const res = await fetch("http://localhost:8080/guardar-seleccion", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
      user_id: userId,
      red_ids: redIds,
    }),
  });

  if (!res.ok) throw new Error("No se pudo guardar");

  return await res.text();
}
