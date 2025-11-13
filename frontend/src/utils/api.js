export async function getUser() {
  const res = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/api/user_zajuna`);
  if (!res.ok) throw new Error("Error obteniendo usuario");
  return res.json();
}

// export async function getCursosActivos() {
//   const res = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/api/cursos_activos`);
//   if (!res.ok) throw new Error("Error obteniendo cursos activos");
//   return res.json();
// }
