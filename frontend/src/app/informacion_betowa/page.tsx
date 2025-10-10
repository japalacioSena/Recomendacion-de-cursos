import InformacionBetowa from "./components/InformacionBetowa";

export default function AnunciosPage() {
  const anuncios = [
    {
      fecha: "09/10/2025",
      titulo: "Nuevo curso de cocina internacional",
      descripcion:
        "A partir del próximo mes abriremos inscripciones para el curso de cocina internacional. Los interesados pueden registrarse en la plataforma.",
    },
    {
      fecha: "15/10/2025",
      titulo: "Mantenimiento del sistema",
      descripcion:
        "El sistema estará en mantenimiento el sábado de 2:00 a 6:00 a.m. Por favor, planifica tus actividades con anticipación.",
    },
  ];
// Renderizado del componente
  return (
    <div className="p-8 bg-gray-50 min-h-screen">
      <h1 className="text-3xl font-bold text-gray-800 mb-6 text-center">
        Anuncios Importantes
      </h1>

      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        {anuncios.map((item, index) => (
          <InformacionBetowa
            key={index}
            fecha={item.fecha}
            titulo={item.titulo}
            descripcion={item.descripcion}
          />
        ))}
      </div>
    </div>
  );
}