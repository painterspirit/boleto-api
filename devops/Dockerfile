# build stage
FROM golang:alpine AS build-env
WORKDIR /go/src/github.com/mundipagg/boleto-api
COPY . .
RUN go build -o boleto-api


FROM alpine:3.7
WORKDIR /home/mundipagg
COPY --from=build-env /go/src/github.com/mundipagg/boleto-api/boleto-api /home/mundipagg/

RUN mkdir -p "/home/mundipagg/boleto_ssh"
RUN mkdir -p "/home/mundipagg/boleto_cert"
RUN mkdir -p "/home/mundipagg/boleto"
RUN mkdir  -p "/home/upMongo"

COPY devops/time "/home/mundipagg/lib/time"
COPY boleto/favicon.ico  "/home/mundipagg/boleto/"
COPY boleto/Arial.ttf "/home/mundipagg/boleto/"

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
RUN chmod +x /home/mundipagg/boleto-api

ENV GOROOT="/home/mundipagg/"
ENV INFLUXDB_HOST="http://localhost"
ENV INFLUXDB_PORT="8086"
ENV PDF_API="http://localhost:7070/topdf"
ENV API_PORT="3000"
ENV API_VERSION="0.0.1"
ENV ENVIROMENT="Development"
ENV SEQ_URL="http://localhost:5341"
ENV SEQ_API_KEY="4jZzTybZ9bUHtJiPdh6"
ENV ENABLE_REQUEST_LOG="false"
ENV ENABLE_PRINT_REQUEST="true"
ENV URL_BB_REGISTER_BOLETO="https://cobranca.homologa.bb.com.br:7101/registrarBoleto"
ENV URL_BB_TOKEN="https://oauth.hm.bb.com.br/oauth/token"
ENV CAIXA_ENV="SGCBS01D"
ENV URL_CAIXA="https://des.barramento.caixa.gov.br/sibar/ManutencaoCobrancaBancaria/Boleto/Externo"
ENV URL_CITI="https://citigroupsoauat.citigroup.com/comercioeletronico/registerboleto/RegisterBoletoSOAP"
ENV URL_CITI_BOLETO="https://ebillpayer.uat.brazil.citigroup.com/ebillpayer/jspInformaDadosConsulta.jsp"
ENV APP_URL="http://localhost:3000/boleto"
ENV ELASTIC_URL="http://localhost:9200"
ENV MONGODB_URL="http://localhost:27017"
ENV MONGODB_USER=""
ENV MONGODB_PASSWORD=""
ENV BOLETO_JSON_STORE="/home/upMongo"
ENV CERT_BOLETO_CRT="/home/mundipagg/boleto_cert/certificate.crt"
ENV CERT_BOLETO_KEY="/home/mundipagg/boleto_cert/pkey.key"
ENV CERT_BOLETO_CA="/home/mundipagg/boleto_cert/ca-cert.ca"
ENV CERT_ICP_BOLETO_KEY="/home/mundipagg/boleto_cert/ICP_PKey.key"
ENV CERT_ICP_BOLETO_CHAIN_CA="/home/mundipagg/boleto_cert/ICP_cadeiaCerts.pem"
ENV URL_SANTANDER_TICKET="https://ymbdlb.santander.com.br/dl-ticket-services/TicketEndpointService"
ENV URL_SANTANDER_REGISTER="https://ymbcash.santander.com.br/ymbsrv/CobrancaEndpointService"
ENV URL_BRADESCO_SHOPFACIL="https://homolog.meiosdepagamentobradesco.com.br/apiboleto/transacao"
ENV ITAU_ENV="1"
ENV SANTANDER_ENV="T"
ENV URL_ITAU_REGISTER="https://gerador-boletos.itau.com.br/router-gateway-app/public/codigo_barras/registro"
ENV URL_ITAU_TICKET="https://oauth.itau.com.br/identity/connect/token"
ENV URL_BRADESCO_NET_EMPRESA="https://cobranca.bradesconetempresa.b.br/ibpjregistrotitulows/registrotitulohomologacao"

ENTRYPOINT ["/home/mundipagg/boleto-api"]
EXPOSE 3000
