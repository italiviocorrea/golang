FROM golang:latest AS build

# Cria o diretorio da aplicacao
RUN mkdir -p $GOPATH/src/github.com/italiviocorrea/golang/grpc/paises

# Definir o diretório de trabalho atual dentro do contêiner
WORKDIR $GOPATH/src/github.com/italiviocorrea/golang/grpc/paises

# Copie tudo do diretório atual para o PWD (Diretório de Trabalho) dentro do contêiner
COPY . .

# Ativar Go Modules
ENV GO111MODULE=on

# Constroi a aplicacao
RUN go build paises.go

# Imagem final
FROM icorrea/alpine:3.9

# Cria um diretorio para aplicacao
RUN mkdir /app

# Copia do container build o aplicativo compilado para diretorio /app
COPY --from=build /go/src/github.com/italiviocorrea/golang/grpc/paises/paises /app

# define o diretorio de trabalho
WORKDIR /app

EXPOSE 50051

CMD ["./paises"]
