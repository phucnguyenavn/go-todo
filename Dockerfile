FROM golang:1.17-alpine

RUN mkdir /app
WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o /todo

EXPOSE 8080

CMD [ "/todo" ]


