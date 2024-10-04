# Use a imagem base do golang
FROM golang:alpine as builder

# Instale git e outras dependências necessárias
RUN apk update && apk add --no-cache git

# Defina o diretório de trabalho dentro do contêiner
WORKDIR /app

# Copie os arquivos go.mod e go.sum
COPY go.mod go.sum ./

# Baixe todas as dependências
RUN go mod download

# Copie o código-fonte para o diretório de trabalho
COPY . .

# Compile o aplicativo Go
RUN go build -o main cmd/api/main.go

# Inicie uma nova etapa a partir de uma imagem base mínima
FROM alpine:latest
RUN apk --no-cache add ca-certificates

# Defina o diretório de trabalho
WORKDIR /root/

# Copie o binário pré-compilado da etapa anterior
COPY --from=builder /app/main .
COPY --from=builder /app/.env .

# Exponha a porta 8008 para o mundo exterior
EXPOSE 8008

# Comando para executar o binário
CMD ["./main"]