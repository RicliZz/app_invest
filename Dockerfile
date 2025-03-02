FROM golang:1.23

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN go build -v -o bin/invest_app ./cmd/app

RUN ls -l bin/
RUN chmod +x bin/invest_app
