FROM golang:1.20-alpine
LABEL org.opencontainers.image.source https://github.com/amundlrohne/dcs-medication-sharing


WORKDIR /app 

COPY go.mod ./
COPY go.sum ./

RUN go install -mod=mod github.com/githubnemo/CompileDaemon
RUN go mod download

COPY *.go ./

#RUN go build -o /consent

EXPOSE 8180

#CMD ["/consent"]
CMD CompileDaemon --build="go build -o /consent" --command=/consent
