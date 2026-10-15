# ---------- etapa 1: compilar ----------
FROM golang:1.26-alpine AS build
WORKDIR /app

# Copia só go.mod/go.sum primeiro: se as dependências não mudarem,
# o Docker reaproveita o cache dessa camada e não baixa tudo de novo.
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o /api .

# ---------- etapa 2: rodar ----------
FROM alpine:3.20
COPY --from=build /api /api
EXPOSE 8080
CMD ["/api"]
