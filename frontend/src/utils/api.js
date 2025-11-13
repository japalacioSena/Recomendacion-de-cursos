export async function getUser() {
  const res = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/api/user_zajuna`);
  if (!res.ok) throw new Error("Error obteniendo usuario");
  return res.json();
}

export async function getTechnologicalRed() {
  const res = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/api/technological_reds`);
  if (!res.ok) throw new Error("Error obteniendo redes tecnológicas");
  return res.json();
}
