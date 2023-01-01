FROM golang:alpine

WORKDIR /app
COPY populate.go go.mod go.sum schema.sql views.sql .

RUN go get -u
RUN go build -o /app/populate

CMD "/app/populate"
