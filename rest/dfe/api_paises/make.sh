#!/usr/bin/env bash

# Constroi a api
go build main.go

# Para e remove os container em execução
docker-compose stop
docker-compose rm -f

# remove a image
docker rmi encat_paises:1.0 -f

# Constroi a imagem docker
docker build -t encat_paises:1.0 .

# Executa os containers conforme docker-compose.yml
docker-compose up -d