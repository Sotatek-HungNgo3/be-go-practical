FROM golang:1.22

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod ./
COPY . .
RUN go mod tidy
RUN go mod download

ENV GIN_MODE release

RUN go build -o payment .

CMD ["./payment"]