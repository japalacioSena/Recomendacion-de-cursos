import ListaDistribucion from "./components/listaDistribución";

export default function CursosActivosPage() {
  return (
    <div className="p-6">
      <h1 className="mt-8 text-2xl font-medium text-gray-900">
        Lista de distribución de cursos
      </h1>
      <div className="bg-gray-200 bg-opacity-25 rounded-lg p-6 mt-4">
        <ListaDistribucion />
      </div>
    </div>
  );
}