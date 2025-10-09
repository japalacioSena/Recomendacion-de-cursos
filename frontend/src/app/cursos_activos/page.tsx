import FilterCursosActivos from "./components/FilterCursosActivos";

export default function CursosActivosPage() {
  return (
    <div className="p-6">
      <h1 className="mt-8 text-2xl font-medium text-gray-900">
        Cursos activos en Betowa
      </h1>
      <div className="bg-gray-200 bg-opacity-25 rounded-lg p-6 mt-4">
        <h2 className="text-xl font-semibold mb-4">Filtro</h2>
        <FilterCursosActivos />
      </div>
    </div>
  );
}

