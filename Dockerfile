FROM golang:1.21-alpine as builder

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o runners-app main.go

EXPOSE 8080
EXPOSE 9000

FROM alpine

WORKDIR /app

COPY --from=builder /app/runners-app /app/runners-app
COPY --from=builder /app/runners.toml /app/runners.toml

CMD [ "./runners-app" ]