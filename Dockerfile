FROM golang:1.20-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o mev-bot ./cmd/bot

FROM alpine:3.18
WORKDIR /app
COPY --from=build /app/mev-bot ./
ENTRYPOINT ["./mev-bot"]
