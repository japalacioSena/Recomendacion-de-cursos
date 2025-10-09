"use client";

type AnuncioProps = {
  fecha: string;
  titulo: string;
  descripcion: string;
};

export default function AnuncioCard({ fecha, titulo, descripcion }: AnuncioProps) {
  return (
    <div className="bg-white shadow-md rounded-xl p-6 w-full max-w-md mx-auto border border-gray-200 hover:shadow-lg transition">
      {/* Fecha */}
      <p className="text-sm text-gray-500 mb-2 text-right">{fecha}</p>

      {/* Título / Anuncio */}
      <h2 className="text-xl font-bold text-gray-900 mb-2">{titulo}</h2>

      {/* Descripción */}
      <p className="text-gray-700 text-justify">{descripcion}</p>
    </div>
  );
}