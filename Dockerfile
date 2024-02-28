FROM golang:1.17-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /app/app

FROM alpine:latest

WORKDIR /app

COPY --from=build /app/app .

EXPOSE 4000 4001

CMD ["./app"]
