FROM golang:1.22.3-alpine as golang

WORKDIR /app

COPY . .

RUN go mod download
RUN go mod verify
RUN go build cmd/main.go

FROM gcr.io/distroless/static-debian11

WORKDIR /app

COPY --from=golang /app/main .

EXPOSE 8080

CMD ["/app/main"]