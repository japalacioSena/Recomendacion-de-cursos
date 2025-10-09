export default function Navbar() {
  return (
    <div className="bg-alternate-background border-b border-gray-100 bg-gray-700">
      <div className="mx-auto px-4 sm:px-6 lg:px-8 flex justify-between h-16">
        <div className="w-full flex items-center gap-2">
          <img src="/images/logo_zanjuna.png" alt="zanjuna" className="w-24" />
          <span className="text-xs font-bold text-white px-2 py-0.5 rounded-lg">
            V1.1.0
          </span>
        </div>

        <div className="flex w-full justify-center items-center gap-2">
          <div id="navbar" className="hidden sm:-my-px gap-3 sm:flex text-white">
            <a href="/">Inicio</a>
            <a href="/cursos_activos">Cursos activos en Betowa</a>
            <a href="">Selección cursos de seguimiento</a>
            <a href="">Selección de cursos de distribución</a>
            <a href="">Información de interés Betowa</a>
          </div>
        </div>

        <div className="w-full flex justify-end items-center gap-2"></div>
      </div>
    </div>
  );
}