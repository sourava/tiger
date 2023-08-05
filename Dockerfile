FROM golang:1.19

COPY . /tiger

WORKDIR /tiger/app

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o ./tiger

EXPOSE 8080

CMD ["./tiger"]