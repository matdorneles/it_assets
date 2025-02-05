FROM golang:1.23.4-alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o backend .

RUN chmod +x /app/backend

# build tiny docker image
FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/backend /app

CMD ["/app/backend"]