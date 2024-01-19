FROM golang:latest as builder
WORKDIR /build
COPY go.mod .
RUN go mod download
COPY . .
ENV CGO_ENABLED=0
ENV GOOS=linux
COPY internal/config/.env /main/.env
RUN go build -o /main/build cmd/api/main.go



FROM alpine:latest

WORKDIR /bin

COPY --from=builder /main/build ./app
COPY --from=builder /main/.env ./.env


ENTRYPOINT ["app"]