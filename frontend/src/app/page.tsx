import Image from "next/image";
import Link from "next/link";

export default function Home() {
  return (
    <div className="p-6">
      <h1 className="mt-8 text-2xl font-medium text-gray-900">Inicio</h1>
      <div className="items-center  bg-gray-200 bg-opacity-25 grid grid-cols-1 md:grid-cols-2 gap-6 lg:gap-8 p-6 lg:p-8">

        {/* Enlace para cursos activos */}
        <Link
          href="/cursos_activos"
          className="inline-flex items-center font-semibold text-indigo-700"
        >
          <div>
            <div className="flex items-center">
              <svg
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 24 24"
                strokeWidth="1.5"
                className="w-6 h-6 stroke-gray-400"
              >
                <path
                  strokeLinecap="round"
                  strokeLinejoin="round"
                  d="M12 6.042A8.967 8.967 0 006 3.75c-1.052 0-2.062.18-3 .512v14.25A8.987 8.987 0 016 18c2.305 0 4.408.867 6 2.292m0-14.25a8.966 8.966 0 016-2.292c1.052 0 2.062.18 3 .512v14.25A8.987 8.987 0 0018 18a8.967 8.967 0 00-6 2.292m0-14.25v14.25"
                />
              </svg>
              <h2 className="ms-3 text-xl font-semibold text-gray-900">
                Cursos activos en Betowa
              </h2>
            </div>

            <p className="mt-4 text-gray-500 text-sm leading-relaxed">
              Accede a la vista de los cursos activos o que estan por salir y
              que son de tu interes para que puedas ingresar a betowa e
              inscribirte.
            </p>

            <p className="mt-4 text-sm">
              Ver
              <svg
                xmlns="http://www.w3.org/2000/svg"
                viewBox="0 0 20 20"
                className="ms-1 w-5 h-5 fill-indigo-500"
              >
                <path
                  fillRule="evenodd"
                  d="M5 10a.75.75 0 01.75-.75h6.638L10.23 7.29a.75.75 0 111.04-1.08l3.5 3.25a.75.75 0 010 1.08l-3.5 3.25a.75.75 0 11-1.04-1.08l2.158-1.96H5.75A.75.75 0 015 10z"
                  clipRule="evenodd"
                />
              </svg>
            </p>
          </div>
        </Link>

        {/* Enlace para seguiemiento de cursos */}
        <Link
          href="/seguimiento_cursos"
          className="inline-flex items-center font-semibold text-indigo-700"
        >
          <div>
            <div className="flex items-center">
              <svg
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 24 24"
                strokeWidth="1.5"
                className="w-6 h-6 stroke-gray-400"
              >
                <path
                  strokeLinecap="round"
                  d="M15.75 10.5l4.72-4.72a.75.75 0 011.28.53v11.38a.75.75 0 01-1.28.53l-4.72-4.72M4.5 18.75h9a2.25 2.25 0 002.25-2.25v-9a2.25 2.25 0 00-2.25-2.25h-9A2.25 2.25 0 002.25 7.5v9a2.25 2.25 0 002.25 2.25z"
                />
              </svg>
              <h2 className="ms-3 text-xl font-semibold text-gray-900">
                Selección seguimiento de cursos
              </h2>
            </div>

            <p className="mt-4 text-gray-500 text-sm leading-relaxed">
              Este módulo permite la selección y seguimiento de cursos o areás
              de interés, en los que esten interesados los usuarios. Para cuando
              hablan los cursos, se les notifique a los usuarios.
            </p>

            <p className="mt-4 text-sm">
              Ver
              <svg
                xmlns="http://www.w3.org/2000/svg"
                viewBox="0 0 20 20"
                className="ms-1 w-5 h-5 fill-indigo-500"
              >
                <path
                  fillRule="evenodd"
                  d="M5 10a.75.75 0 01.75-.75h6.638L10.23 7.29a.75.75 0 111.04-1.08l3.5 3.25a.75.75 0 010 1.08l-3.5 3.25a.75.75 0 11-1.04-1.08l2.158-1.96H5.75A.75.75 0 015 10z"
                  clipRule="evenodd"
                />
              </svg>
            </p>
          </div>
        </Link>

        {/* Enlace para curso de distribución */}
        <Link
          href="/lista_distribucion"
          className="inline-flex items-center font-semibold text-indigo-700"
        >
          <div>
            <div className="flex items-center">
              <svg
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 24 24"
                strokeWidth="1.5"
                className="w-6 h-6 stroke-gray-400"
              >
                <path
                  strokeLinecap="round"
                  strokeLinejoin="round"
                  d="M2.25 15.75l5.159-5.159a2.25 2.25 0 013.182 0l5.159 5.159m-1.5-1.5l1.409-1.409a2.25 2.25 0 013.182 0l2.909 2.909m-18 3.75h16.5a1.5 1.5 0 001.5-1.5V6a1.5 1.5 0 00-1.5-1.5H3.75A1.5 1.5 0 002.25 6v12a1.5 1.5 0 001.5 1.5zm10.5-11.25h.008v.008h-.008V8.25zm.375 0a.375.375 0 11-.75 0 .375.375 0 01.75 0z"
                />
              </svg>
              <h2 className="ms-3 text-xl font-semibold text-gray-900">
                Selección de cursos de distribución
              </h2>
            </div>

            <p className="mt-4 text-gray-500 text-sm leading-relaxed">
              Este módulo permite dejar la lista de cursos que te recomienda el
              sistema, de acuerdo a los cursos que haz realizado. Puedes
              eliminar o agregar cursos a esta lista.
            </p>
            <p className="mt-4 text-sm">
              Ver
              <svg
                xmlns="http://www.w3.org/2000/svg"
                viewBox="0 0 20 20"
                className="ms-1 w-5 h-5 fill-indigo-500"
              >
                <path
                  fillRule="evenodd"
                  d="M5 10a.75.75 0 01.75-.75h6.638L10.23 7.29a.75.75 0 111.04-1.08l3.5 3.25a.75.75 0 010 1.08l-3.5 3.25a.75.75 0 11-1.04-1.08l2.158-1.96H5.75A.75.75 0 015 10z"
                  clipRule="evenodd"
                />
              </svg>
            </p>
          </div>
        </Link>

        {/* Información de interes betowa */}
        <Link
          href="/informacion_betowa"
          className="inline-flex items-center font-semibold text-indigo-700"
        >
          <div>
            <div className="flex items-center">
              <svg
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 24 24"
                strokeWidth="1.5"
                className="w-6 h-6 stroke-gray-400"
              >
                <path
                  strokeLinecap="round"
                  strokeLinejoin="round"
                  d="M16.5 10.5V6.75a4.5 4.5 0 10-9 0v3.75m-.75 11.25h10.5a2.25 2.25 0 002.25-2.25v-6.75a2.25 2.25 0 00-2.25-2.25H6.75a2.25 2.25 0 00-2.25 2.25v6.75a2.25 2.25 0 002.25 2.25z"
                />
              </svg>
              <h2 className="ms-3 text-xl font-semibold text-gray-900">
                Información de interés Betowa
              </h2>
            </div>
            <p className="mt-4 text-gray-500 text-sm leading-relaxed">
              Este módulo se encarga de mostrar información relevante sobre
              Betowa, como noticias, actualizaciones, eventos próximos y otros
              datos que puedan ser de interés para los usuarios.
            </p>
            <p className="mt-4 text-sm">
              Ver
              <svg
                xmlns="http://www.w3.org/2000/svg"
                viewBox="0 0 20 20"
                className="ms-1 w-5 h-5 fill-indigo-500"
              >
                <path
                  fillRule="evenodd"
                  d="M5 10a.75.75 0 01.75-.75h6.638L10.23 7.29a.75.75 0 111.04-1.08l3.5 3.25a.75.75 0 010 1.08l-3.5 3.25a.75.75 0 11-1.04-1.08l2.158-1.96H5.75A.75.75 0 015 10z"
                  clipRule="evenodd"
                />
              </svg>
            </p>
          </div>
        </Link>
      </div>
    </div>
  );
}
