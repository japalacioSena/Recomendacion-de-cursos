"use client";
import { useEffect, useState } from "react";
import Link from "next/link";
import { usePathname } from "next/navigation"; // ✅ Import correcto
import { getUser } from "@/utils/api";

const links = [
  { href: "/", name: "Inicio" },
  { href: "/cursos_activos", name: "Cursos activos" },
  { href: "/seguimiento_cursos", name: "Seguimiento" },
  { href: "/lista_distribucion", name: "Lista de distribución" },
  { href: "/informacion_betowa", name: "Información Betowa" },
];

export default function Navbar() {
  const [user, setUser] = useState(null);
  const pathname = usePathname(); // ✅ Hook que obtiene la ruta actual

  useEffect(() => {
    getUser()
      .then((data) => {
        console.log("✅ Usuario recibido:", data);
        setUser(data);
      })
      .catch((err) => console.error("❌ Error al obtener el usuario:", err));
  }, []);

  return (
    <div className="bg-gray-700 border-b border-gray-100">
      <div className="mx-auto px-4 sm:px-6 lg:px-8 flex justify-between h-16">
        {/* LOGO */}
        <div className="w-full flex items-center gap-2">
          <img src="/images/logo_zanjuna.png" alt="zanjuna" className="w-24" />
          <span className="text-xs font-bold text-white px-2 py-0.5 rounded-lg">
            V1.1.0
          </span>
        </div>

        {/* NAV LINKS */}
        <div className="flex w-full justify-center items-center gap-4">
          {links.map((link) => {
            // ✅ Define isActive dentro del map
            const isActive = pathname === link.href;
            return (
              <Link
                key={link.name}
                href={link.href}
                className={`inline-flex items-center px-3 py-1 border-b-2 text-sm font-medium transition 
                  ${
                    isActive
                      ? "border-white text-white" // activo
                      : "border-transparent text-gray-300 hover:border-white hover:text-gray-100" // inactivo
                  }`}
              >
                {link.name}
              </Link>
            );
          })}
        </div>

        {/* USER INFO */}
        <div className="w-full flex justify-end items-center gap-2 text-white font-semibold">
          {user ? `${user.firstname} ${user.lastname}` : "Cargando..."}
        </div>
      </div>
    </div>
  );
}
