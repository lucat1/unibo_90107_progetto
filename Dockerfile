FROM golang:alpine

WORKDIR /app
COPY . .

RUN go get -u
RUN go build -o /app/populate

CMD "/app/populate"
