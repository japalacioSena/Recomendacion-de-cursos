"use client";

import { useState, useMemo } from "react";

type Interes = {
  id: number;
  nombre: string;
};

export default function TablaSeguimiento() {
  // 🔹 Datos quemados simulando respuesta de un backend Go
  const intereses: Interes[] = [
    { id: 1, nombre: "Cocina" },
    { id: 2, nombre: "Base de Datos" },
    { id: 3, nombre: "Inglés" },
    { id: 4, nombre: "Videojuegos" },
    { id: 5, nombre: "Programación" },
    { id: 6, nombre: "Diseño Gráfico" },
    { id: 7, nombre: "Matemáticas" },
    { id: 8, nombre: "Historia" },
    { id: 9, nombre: "Física" },
  ];

  // 🔹 Estado para filtro y checkboxes seleccionados
  const [busqueda, setBusqueda] = useState("");
  const [seleccionados, setSeleccionados] = useState<number[]>([]);
  const [guardando, setGuardando] = useState(false);

  // 🔹 Filtrar según búsqueda (insensible a mayúsculas)
  const resultados = useMemo(() => {
    return intereses.filter((item) =>
      item.nombre.toLowerCase().includes(busqueda.toLowerCase())
    );
  }, [busqueda]);

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
   // 🔹 Guardar selección (por ahora muestra en consola)
  const handleGuardar = async () => {
    setGuardando(true);

    // Simulamos envío al backend Go
    const seleccionadosData = intereses.filter((i) =>
      seleccionados.includes(i.id)
    );

    console.log("Datos a guardar:", seleccionadosData);

    // Ejemplo de cómo será con backend Go:
    /*
    await fetch("http://localhost:8080/api/intereses/guardar", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(seleccionadosData),
    });
    */

    // Simulamos retardo
    await new Promise((r) => setTimeout(r, 1000));
    setGuardando(false);
    alert("¡Intereses guardados correctamente!");
  };

  
  return (
    <div className="p-6">
      <h2 className="text-2xl font-bold mb-4">Tabla de Intereses</h2>

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
              {/* Si la última fila no tiene 3 celdas, completar vacíos */}
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
