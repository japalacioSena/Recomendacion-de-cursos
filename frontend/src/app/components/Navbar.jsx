"use client";

import Link from "next/link";

const links = [
  { href: "/", name: "Inicio" },
  { href: "/cursos_activos", name: "Cursos activos en Betowa" },
  { href: "/seguimiento_cursos", name: "Seguimiento de cursos" },
  { href: "/lista_distribucion", name: "Lista de distribución" },
  { href: "/informacion_betowa", name: "Información de interés Betowa" },
];

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
          <div
            id="navbar"
            className="hidden sm:-my-px gap-3 sm:flex text-white"
          >
            {links.map((link) => (
              <Link
                key={link.name}
                href={link.href}
                className="inline-flex items-center px-1 pt-1 border-b-2 border-transparent text-sm font-medium leading-5 hover:border-white hover:text-gray-300 focus:outline-none focus:border-white transition"
              >
                {link.name}
              </Link>
            ))}
          </div>
        </div>

        <div className="w-full flex justify-end items-center gap-2"></div>
      </div>
    </div>
  );
}
