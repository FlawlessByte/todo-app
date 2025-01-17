FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o /todo-app cmd/main.go

EXPOSE 8080

CMD ["/todo-app"]
