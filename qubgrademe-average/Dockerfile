FROM golang:1.17

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY src/*.go ./

RUN go build -o /qubgrademe-average

EXPOSE 1323

CMD [ "/qubgrademe-average" ]