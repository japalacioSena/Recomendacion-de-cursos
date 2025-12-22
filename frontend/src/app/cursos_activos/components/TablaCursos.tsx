"use client";
import React from "react";

// La estructura de la data que recibes del backend de Go
export interface Programa {
  id: number;
  name: string;
  registration_date: string;
  active_date: string;
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

  return (
    <div className="overflow-x-auto mt-6">
      <table className="min-w-full border border-gray-300 rounded-lg shadow-md">
        <thead className="bg-gray-200">
          <tr>
            {/* 🚨 Actualizamos los encabezados para coincidir con la data */}
            <th className="px-4 py-2 border text-left">ID (Código)</th>
            <th className="px-4 py-2 border text-left">Nombre del Curso</th>
            <th className="px-4 py-2 border text-left">Fecha de Registro</th>
            <th className="px-4 py-2 border text-left">Fecha de Actividad</th>
            {/* Si el estado no viene del backend, puedes quitar esta columna */}
            <th className="px-4 py-2 border text-left">Estado</th> 
          </tr>
        </thead>
        <tbody>
          {data.map((programa, index) => (
            <tr key={index} className="hover:bg-gray-50">
              {/* 🚨 Usamos los nombres de propiedades del JSON de Go */}
              <td className="px-4 py-2 border">{programa.id}</td>
              <td className="px-4 py-2 border">{programa.name}</td>
              <td className="px-4 py-2 border">{programa.registration_date}</td>
              <td className="px-4 py-2 border">{programa.active_date}</td>
              {/* Rellena con un valor estático si el backend no lo provee */}
              <td className="px-4 py-2 border">Activo</td> 
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default TablaCursos;