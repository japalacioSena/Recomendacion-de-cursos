import TablaSeguimiento from "./components/TablaSeguimiento";

export default function SeguimientoCursosPage() {
  return (
    <div className="p-6">
      <h1 className="mt-8 text-2xl font-medium text-gray-900">
        Seguimiento de cursos
      </h1>
      <div className="bg-gray-200 bg-opacity-25 rounded-lg p-6 mt-4">
        <TablaSeguimiento />
      </div>
    </div>
  );
}