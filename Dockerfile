FROM golang:1.15-alpine as mother

WORKDIR /src/

COPY theo.go /src/
RUN apk add git --update && rm -rf /var/cache/apk/*
RUN go get github.com/gorilla/mux
RUN CGO_ENABLED=0 go build -o /opt/taas

FROM alpine:3.14
COPY --from=mother /opt/taas /opt/taas
COPY insults.json /opt/insults.json
EXPOSE 8080
WORKDIR /opt/
CMD ["./taas"]
