FROM golang:1.20-alpine
LABEL org.opencontainers.image.source https://github.com/amundlrohne/dcs-medication-sharing

WORKDIR /app 

COPY go.mod ./
COPY go.sum ./

RUN go install -mod=mod github.com/githubnemo/CompileDaemon
RUN go mod download

COPY *.go ./

#RUN go build -o /medication-record

EXPOSE 8083

#CMD [ "/medication-record" ]
CMD CompileDaemon --build="go build -o /medication-record" --command=/medication-record
