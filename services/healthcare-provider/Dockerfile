FROM golang:1.20-alpine
LABEL org.opencontainers.image.source https://github.com/amundlrohne/dcs-medication-sharing

WORKDIR /app 

COPY go.mod ./
COPY go.sum ./

RUN go install -mod=mod github.com/githubnemo/CompileDaemon
RUN go mod download

COPY *.go ./

#RUN go build -o /healthcare-provider

EXPOSE 8280

#CMD [ "/healthcare-provider" ]
CMD CompileDaemon --build="go build -o /healthcare-provider" --command=/healthcare-provider