FROM golang:1.25.1 AS builder

WORKDIR /app
# Cache deps first
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /out/server ./cmd/server

# --- Runtime stage ---
FROM gcr.io/distroless/base-debian12

WORKDIR /app
COPY --from=builder /out/server /app/server

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/app/server"]
