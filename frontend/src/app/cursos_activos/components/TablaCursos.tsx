"use client";
import React from "react";

// Definición del tipo Programa
export interface Programa {
  codigo: string;
  nombre: string;
  estado: string;
  fecha_apertura: string;
  fecha_cierre: string;
}
// Props para el componente TablaCursos
interface TablaCursosProps {
  data: Programa[];
}

// Componente para mostrar la tabla de programas
const TablaCursos: React.FC<TablaCursosProps> = ({ data }) => {
  if (data.length === 0) {
    return (
      <div className="text-center text-gray-500 mt-6">
        No hay programas que coincidan con el filtro.
      </div>
    );
  }

  // Renderizado de la tabla
  return (
    <div className="overflow-x-auto mt-6">
      <table className="min-w-full border border-gray-300 rounded-lg shadow-md">
        <thead className="bg-gray-200">
          <tr>
            <th className="px-4 py-2 border text-left">Código</th>
            <th className="px-4 py-2 border text-left">Nombre</th>
            <th className="px-4 py-2 border text-left">Estado</th>
            <th className="px-4 py-2 border text-left">Fecha de apertura</th>
            <th className="px-4 py-2 border text-left">Fecha de cierre</th>
          </tr>
        </thead>
        <tbody>
          {data.map((programa, index) => (
            <tr key={index} className="hover:bg-gray-50">
              <td className="px-4 py-2 border">{programa.codigo}</td>
              <td className="px-4 py-2 border">{programa.nombre}</td>
              <td className="px-4 py-2 border">{programa.estado}</td>
              <td className="px-4 py-2 border">{programa.fecha_apertura}</td>
              <td className="px-4 py-2 border">{programa.fecha_cierre}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default TablaCursos;
