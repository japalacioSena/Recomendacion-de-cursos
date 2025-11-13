"use client";

import { useState, useMemo, useEffect } from "react";
import { getTechnologicalRed } from "@/utils/api";


type Interes = {
  id: number;
  nombre: string;
};

export default function TablaSeguimiento() {
  // 🔹 Estado para los datos del backend
  const [intereses, setIntereses] = useState<Interes[]>([]);
  const [cargando, setCargando] = useState(true);
  const [error, setError] = useState<string | null>(null);

  // 🔹 Estado para búsqueda y selección
  const [busqueda, setBusqueda] = useState("");
  const [seleccionados, setSeleccionados] = useState<number[]>([]);
  const [guardando, setGuardando] = useState(false);

  // 🔹 Cargar datos al montar el componente
  useEffect(() => {
    async function fetchData() {
      try {
        const data = await getTechnologicalRed();
        // 🔸 Ajustar nombres del backend (Go usa "name")
        const mapped = data.map((item: any) => ({
          id: item.id,
          nombre: item.name,
        }));
        setIntereses(mapped);
      } catch (err) {
        console.error(err);
        setError("No se pudieron obtener las redes tecnológicas");
      } finally {
        setCargando(false);
      }
    }
    fetchData();
  }, []);

  // 🔹 Filtrar según búsqueda
  const resultados = useMemo(() => {
    return intereses.filter((item) =>
      item.nombre.toLowerCase().includes(busqueda.toLowerCase())
    );
  }, [busqueda, intereses]);

  // 🔹 Alternar selección
  const toggleSeleccion = (id: number) => {
    setSeleccionados((prev) =>
      prev.includes(id) ? prev.filter((i) => i !== id) : [...prev, id]
    );
  };

  // 🔹 Dividir resultados en filas de 3 columnas
  const filas = [];
  for (let i = 0; i < resultados.length; i += 3) {
    filas.push(resultados.slice(i, i + 3));
  }

  // 🔹 Guardar selección
  const handleGuardar = async () => {
    setGuardando(true);
    const seleccionadosData = intereses.filter((i) =>
      seleccionados.includes(i.id)
    );

    console.log("📦 Datos a guardar:", seleccionadosData);

    // Aquí podrías hacer un POST al backend si quisieras guardar
    await new Promise((r) => setTimeout(r, 1000));

    setGuardando(false);
    alert("¡Intereses guardados correctamente!");
  };

  // 🔹 Mostrar estados
  if (cargando) return <p className="p-4">Cargando redes tecnológicas...</p>;
  if (error) return <p className="p-4 text-red-500">{error}</p>;

  return (
    <div className="p-6">
      <h2 className="text-2xl font-bold mb-4">Redes Tecnológicas</h2>

      {/* 🔍 Campo de búsqueda */}
      <input
        type="text"
        placeholder="Buscar por nombre..."
        className="border rounded-lg p-2 mb-4 w-full md:w-1/2 focus:ring focus:ring-blue-300"
        value={busqueda}
        onChange={(e) => setBusqueda(e.target.value)}
      />

      {/* 📋 Tabla */}
      <table className="min-w-full border border-gray-300 rounded-lg">
        <tbody>
          {filas.map((fila, index) => (
            <tr key={index} className="border-b border-gray-200">
              {fila.map((item) => (
                <td key={item.id} className="p-4 text-center border-r">
                  <label className="flex items-center justify-center gap-2 cursor-pointer">
                    <input
                      type="checkbox"
                      checked={seleccionados.includes(item.id)}
                      onChange={() => toggleSeleccion(item.id)}
                      className="h-5 w-5 text-blue-600 rounded"
                    />
                    <span>{item.nombre}</span>
                  </label>
                </td>
              ))}
              {fila.length < 3 &&
                Array.from({ length: 3 - fila.length }).map((_, i) => (
                  <td key={`vacio-${i}`} className="p-4 border-r"></td>
                ))}
            </tr>
          ))}
        </tbody>
      </table>

      {/* 📦 Debug / vista previa de seleccionados */}
      <div className="mt-6">
        <h3 className="font-semibold">Seleccionados:</h3>
        <p>
          {seleccionados.length > 0
            ? seleccionados
                .map(
                  (id) => intereses.find((interes) => interes.id === id)?.nombre
                )
                .join(", ")
            : "Ninguno"}
        </p>
      </div>

      {/* 💾 Botón de Guardar */}
      <button
        onClick={handleGuardar}
        disabled={guardando}
        className={`mt-6 px-6 py-2 rounded-lg text-white ${
          guardando ? "bg-gray-400" : "bg-blue-600 hover:bg-blue-700"
        } transition`}
      >
        {guardando ? "Guardando..." : "Guardar"}
      </button>
    </div>
  );
}
