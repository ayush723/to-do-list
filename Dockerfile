FROM golang:1.16-alpine

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o /docker-todo

EXPOSE 8080

CMD ["/docker-todo"]
