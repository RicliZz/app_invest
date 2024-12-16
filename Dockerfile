FROM golang:1.23

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN go build -v -o /usr/src/app/bin/invest_app ./cmd/app

RUN ls -l /usr/src/app/bin/
RUN chmod +x /usr/src/app/bin/invest_app

CMD ["sh", "-c", "until pg_isready -h invest_postgres -U postgres; do echo waiting for postgres; sleep 2; done; /usr/src/app/bin/invest_app"]
