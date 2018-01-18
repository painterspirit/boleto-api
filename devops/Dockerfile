FROM alpine:latest

RUN mkdir -p "/home/mundipagg/"
RUN mkdir -p "/home/mundipagg/boleto_ssh"

ADD boleto-api /home/mundipagg/

COPY time "/home/mundipagg/lib/time"

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
RUN chmod +x /home/mundipagg/boleto-api
RUN ls -la /home/mundipagg/
RUN mkdir  -p "/home/upMongo"
ENTRYPOINT ["/home/mundipagg/boleto-api"]
EXPOSE 3000
