# Imagen base de Go
FROM golang:1.20

# Establecer directorio de trabajo
WORKDIR /app

# Copiar archivos
COPY . .

# Descargar dependencias
RUN go mod tidy

# Compilar el código
RUN go build -o main .

# Exponer el puerto
EXPOSE 2002

# Ejecutar el servicio
CMD ["./main"]
