FROM golang:1.23-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN go build -o /chess-server

EXPOSE 8080

CMD [ "/chess-server" ]