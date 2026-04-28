FROM golang:1.26-alpine AS builder

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags prod -o /out/server ./cmd/server

FROM gcr.io/distroless/static-debian12:nonroot

WORKDIR /app

COPY --from=builder /out/server /app/server
COPY static/config/production/config.toml /app/config.toml
COPY internal/persistence/database/fake-database.json /app/internal/persistence/database/fake-database.json

EXPOSE 3001

ENTRYPOINT ["/app/server"]
