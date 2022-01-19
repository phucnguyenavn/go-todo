FROM golang:${GO_VERSION}

WORKDIR /go-todo/main

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /go-todo/main

EXPOSE 8080

CMD [ "/go-todo/main" ]
