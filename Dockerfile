FROM golang:latest as builder
WORKDIR /build
COPY go.mod .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /main cmd/api/main.go
FROM scratch
COPY --from=builder main /bin/main


# APP
ENV PORT=2000
ENV ENVIRONMENT="development"
ENV DEBUG=true

# DATABASE

# POSTGRE
ENV DB_POSTGRE_DRIVER="postgres"
ENV DB_POSTGRE_DSN="user=postgres password=postgres host=localhost port=5432 dbname=change-it sslmode=disable timezone=Europe/Moscow"
ENV DB_POSTGRE_URL="postgres://user:pass@host/db"


ENTRYPOINT ["/bin/main"]