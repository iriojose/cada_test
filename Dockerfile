# Establece la imagen base de Golang (puedes usar otra versión si lo prefieres)
FROM golang:alpine as builder

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia todo el código fuente de tu aplicación
COPY go.mod go.sum ./
COPY . .

#get air dependencies
#RUN go get -u github.com/cosmtrek/air

# Compila la aplicación en una versión estática
RUN go build -o main .

# Copia el binario compilado desde el primer paso
#COPY --from=builder /app/main .

# Exponemos el puerto 8000
EXPOSE 8000

# Establece el comando de inicio para tu aplicación
CMD ["./main"]


#[default]
#region = us-east-1
#aws_access_key_id=AKIA3DDPM5PQ4WL6HZTL
#aws_secret_access_key=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