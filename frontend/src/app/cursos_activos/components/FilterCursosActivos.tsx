"use client";
import { useState, useEffect } from "react";
import TablaCursos, { Programa } from "./TablaCursos";

const FilterCursosActivos: React.FC = () => {
  const [selectedOption, setSelectedOption] = useState("codigo");
  const [inputValue, setInputValue] = useState("");
  const [data, setData] = useState<Programa[]>([]);
  const [filteredData, setFilteredData] = useState<Programa[]>([]);

  // Datos simulados (en el futuro vendrán de tu backend en Go)
  useEffect(() => {
    const mockData: Programa[] = [
      {
        codigo: "001",
        nombre: "Matemáticas Aplicadas",
        estado: "Activo",
        fecha_apertura: "10/01/2024",
        fecha_cierre: "10/01/2025",
      },
      {
        codigo: "002",
        nombre: "Física Experimental",
        estado: "Inactivo",
        fecha_apertura: "05/03/2023",
        fecha_cierre: "05/03/2024",
      },
      {
        codigo: "003",
        nombre: "Programación Avanzada",
        estado: "Activo",
        fecha_apertura: "01/06/2024",
        fecha_cierre: "01/06/2025",
      },
    ];
    setData(mockData);
    setFilteredData(mockData);
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
          <option value="codigo">Código del programa</option>
          <option value="nombre">Nombre del programa</option>
          <option value="estado">Estado del programa</option>
          <option value="fecha_apertura">Fecha de apertura</option>
          <option value="fecha_cierre">Fecha de cierre</option>
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
