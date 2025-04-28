FROM golang:latest

WORKDIR /app

COPY . .

RUN touch .env

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]