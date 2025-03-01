FROM golang:1.23

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN go build -v -o /usr/src/app/bin/invest_app ./cmd/app

RUN ls -l /usr/src/app/bin/
RUN chmod +x /usr/src/app/bin/invest_app
