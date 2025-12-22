"use client";
import { useState, useEffect } from "react";
import TablaCursos, { Programa } from "./TablaCursos";
import { getUser } from "@/utils/api";

const FilterCursosActivos: React.FC = () => {
  const [selectedOption, setSelectedOption] = useState("codigo");
  const [inputValue, setInputValue] = useState("");
  const [data, setData] = useState<Programa[]>([]);
  const [filteredData, setFilteredData] = useState<Programa[]>([]);

  // Datos simulados (en el futuro vendrán de tu backend en Go)
  useEffect(() => {
    const loadCursos = async () => {
      try {
        // 1️⃣ Obtener usuario
        const usuario = await getUser();

        // 2️⃣ Llamar backend pasando el id_user
        const res = await fetch(
          `http://localhost:8080/api/denomination-cursos?id_user=${usuario.id}`,
          { credentials: "include" }
        );

        if (!res.ok) {
          const text = await res.text();
          throw new Error(text);
        }

        const data = await res.json();
        setData(data);
        setFilteredData(data);
      } catch (err: any) {
        console.error("❌ Error backend:", err.message);
      }
    };

    loadCursos(); // 👈 ejecutar función async
  }, []);

  // Filtrar automáticamente cuando cambia el valor
  useEffect(() => {
    if (!inputValue) {
      setFilteredData(data);
      return;
    }

    const filtered = data.filter((item) => {
      const field = selectedOption as keyof Programa;
      return item[field]
        .toString()
        .toLowerCase()
        .includes(inputValue.toLowerCase());
    });

    setFilteredData(filtered);
  }, [inputValue, selectedOption, data]);

  // Manejar cambio en el select y limpiar input
  const handleSelectChange = (e: React.ChangeEvent<HTMLSelectElement>) => {
    setSelectedOption(e.target.value);
    setInputValue("");
  };

  // Manejar cambio en el input de fecha
  const handleDateChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const date = e.target.value;
    if (!date) return;
    const [year, month, day] = date.split("-");
    const formatted = `${day}/${month}/${year}`;
    setInputValue(formatted);
  };

  const isDateField =
    selectedOption === "fecha_apertura" || selectedOption === "fecha_cierre";

  // Renderizado del componente
  return (
    <div className="w-full max-w-4xl mx-auto mt-10">
      <div className="flex flex-col md:flex-row items-center gap-2">
        {/* Select para filtros */}
        <select
          value={selectedOption}
          onChange={handleSelectChange}
          className="border rounded-lg p-2 h-11 w-full md:w-1/2 focus:ring focus:ring-blue-300 appearance-none"
        >
          <option value="id">Código del curso</option>
          <option value="name">Nombre del curso</option>
          <option value="registration_date">Fecha de registro</option>
          <option value="active_date">Fecha de actividad</option>
        </select>

        {/* Campo dinámico */}
        {isDateField ? (
          <input
            type="date"
            className="border rounded-lg p-2 h-11 w-full md:w-1/2 focus:ring focus:ring-blue-300"
            onChange={handleDateChange}
            value={
              inputValue
                ? (() => {
                    // Si el valor está en formato dd/mm/yyyy, lo convertimos a yyyy-mm-dd para el input date
                    const [day, month, year] = inputValue.split("/");
                    return `${year}-${month}-${day}`;
                  })()
                : ""
            }
          />
        ) : (
          <input
            type="text"
            placeholder="Ingrese un valor"
            value={inputValue ?? ""}
            onChange={(e) => setInputValue(e.target.value)}
            className="border rounded-lg p-2 h-11 w-full md:w-1/2 focus:ring focus:ring-blue-300"
          />
        )}
      </div>

      {/* Tabla de resultados */}
      <TablaCursos data={filteredData} />
    </div>
  );
};

export default FilterCursosActivos;
