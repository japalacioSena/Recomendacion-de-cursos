# Imagen base de Go
FROM golang:alpine

# Establece el directorio de trabajo
WORKDIR /app

# Copia todos los archivos del backend al contenedor
COPY . .

# Descarga las dependencias
RUN go mod tidy

# Compila el binario del servidor
RUN go build -o server .

# Expone el puerto 8080
EXPOSE 8080

# Ejecuta el binario
CMD ["./server"]