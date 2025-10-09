# Usa Node.js basado en Debian (no Alpine)
FROM node:20-slim

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Instala dependencias necesarias del sistema
RUN apt-get update && apt-get install -y python3 build-essential && rm -rf /var/lib/apt/lists/*

# Copia los archivos de dependencias
COPY package*.json ./

# Instala las dependencias de npm
RUN npm install

# Copia el resto del código del proyecto
COPY . .

# Expone el puerto de Next.js
EXPOSE 3000

# Ejecuta el servidor de desarrollo
CMD ["npm", "run", "dev"]