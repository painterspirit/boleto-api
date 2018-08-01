FROM alpine:3.7

RUN mkdir -p "/home/mundipagg/"
RUN mkdir -p "/home/mundipagg/boleto_ssh"
RUN mkdir -p "/home/mundipagg/boleto_cert"
RUN mkdir -p "/home/mundipagg/boleto"

ENV GOROOT="/home/mundipagg/"

COPY time "/home/mundipagg/lib/time"
COPY favicon.ico  "/home/mundipagg/boleto/"
COPY Arial.ttf "/home/mundipagg/boleto/"

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
ADD boleto-api /home/mundipagg/
RUN chmod +x /home/mundipagg/boleto-api
RUN mkdir  -p "/home/upMongo"
WORKDIR /home/mundipagg
ENTRYPOINT ["/home/mundipagg/boleto-api"]
EXPOSE 3000
