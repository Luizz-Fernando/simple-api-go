FROM golang:1.24

# Instala o Air a partir do caminho correto
RUN go install github.com/air-verse/air@latest

# Garante que o binário do Go esteja no PATH
ENV PATH="/go/bin:${PATH}"

WORKDIR /go/src/app

# (Opcional, pois o volume vai sobrescrever, mas bom para produção)
COPY go.mod go.sum ./
RUN go mod download

COPY . .

EXPOSE 8000

# Inicia o Air em modo live-reload
CMD ["air"]