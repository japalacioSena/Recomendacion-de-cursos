"use client";
import { useState } from "react";

export default function ListaDistribucion() {
  const [suscrito, setSuscrito] = useState(false);
  const [loading, setLoading] = useState(false);

  const handleToggle = async () => {
    setLoading(true);

    // 🔹 Simulamos el envío al backend con un retardo de 1.5 segundos
    await new Promise((resolve) => setTimeout(resolve, 1500));

    // 🔹 Cambiamos el estado de suscripción
    setSuscrito(!suscrito);
    setLoading(false);

    console.log(
      !suscrito ? "Simulación: suscripción realizada" : "Simulación: baja realizada"
    );
  };

  return (
    <div className="flex flex-col items-center gap-4 p-6 bg-white rounded-lg shadow-md w-full max-w-md mx-auto">
      <h2 className="text-xl font-bold">Notificaciones del Campus</h2>
      <p className="text-gray-600">
        {suscrito
          ? "Estás suscrito a las notificaciones."
          : "No estás suscrito actualmente."}
      </p>

      <button
        onClick={handleToggle}
        disabled={loading}
        className={`px-4 py-2 rounded text-white transition ${
          suscrito ? "bg-red-600 hover:bg-red-700" : "bg-green-600 hover:bg-green-700"
        } ${loading ? "opacity-70 cursor-not-allowed" : ""}`}
      >
        {loading
          ? "Procesando..."
          : suscrito
          ? "Darse de baja"
          : "Suscribirse"}
      </button>
    </div>
  );
}
