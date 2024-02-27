# ---- Build Phase ----
FROM golang:1.17-alpine AS build

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -o /app/app

FROM alpine:latest

WORKDIR /app

COPY --from=build /app/app .

EXPOSE 8000

CMD ["./app"]
