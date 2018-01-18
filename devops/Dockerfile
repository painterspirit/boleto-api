FROM alpine:3.7

RUN mkdir -p "/home/mundipagg/"
RUN mkdir -p "/home/mundipagg/boleto_ssh"
RUN mkdir -p "/home/mundipagg/boleto_cert"

ENV GOROOT="/home/mundipagg/"

COPY time "/home/mundipagg/lib/time"

ADD boleto-api /home/mundipagg/
RUN chmod +x /home/mundipagg/boleto-api
RUN mkdir  -p "/home/upMongo"
ENTRYPOINT ["/home/mundipagg/boleto-api"]
EXPOSE 3000
