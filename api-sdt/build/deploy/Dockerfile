FROM golang:1-alpine3.16 as builder
WORKDIR /work

# download dos modulos em camada separada para permitir cache para o docker build
COPY ./go.mod ./go.sum ./
RUN go mod download

# copiar o codigo fonte
COPY . .

RUN CGO_ENABLED=0 go build -o apisdt ./cmd/apisdt/main.go

FROM alpine:3.16
MAINTAINER Italivio Correa <italiviocorrea@gmail.com>

RUN apk add --no-cache ca-certificates

WORKDIR /bin
COPY --from=builder /work/apisdt /bin/apisdt
ADD ./api/swaggerui /bin/api/swaggerui
COPY app-prd.env /bin/app.env

# COPY ./app.env ./app.env
# variaveis de ambiente com valores default

EXPOSE 8080

CMD /bin/apisdt
