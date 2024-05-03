FROM golang:1.17

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY .env ./
RUN go mod download

COPY src/*.go ./

RUN go build -o /qubgrademe-monitor

EXPOSE 1326

CMD [ "/qubgrademe-monitor" ]
