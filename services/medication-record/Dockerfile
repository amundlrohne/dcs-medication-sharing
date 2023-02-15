FROM golang:1.20-alpine

WORKDIR /app 

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./

RUN go build -o /medication-record

EXPOSE 8083

CMD [ "/medication-record" ]