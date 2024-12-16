FROM golang:1.22 AS build

WORKDIR /cmd

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/app -v ./cmd/app

FROM alpine:3.20 AS final

WORKDIR /

COPY --from=build /bin/app /app
COPY .env .env

EXPOSE 8080
