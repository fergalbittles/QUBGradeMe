FROM golang:1.17

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY .env ./
RUN go mod download

COPY src/*.go ./

RUN go build -o /qubgrademe-proxy

EXPOSE 1324

CMD [ "/qubgrademe-proxy" ]
