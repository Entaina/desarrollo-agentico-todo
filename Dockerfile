FROM golang:1.26-alpine AS build

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o /out/todo ./cmd/app

FROM alpine:3.22

RUN addgroup -S app && adduser -S app -G app

WORKDIR /app

COPY --from=build /out/todo /usr/local/bin/todo
COPY database.yml .

ENV ADDR=0.0.0.0 \
    PORT=3000

EXPOSE 3000
USER app

CMD ["todo"]
