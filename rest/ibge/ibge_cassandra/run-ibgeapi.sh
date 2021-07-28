#!/usr/bin/env bash

# Constroi a ibgeapims
go build

# Para e remove os container em execução
docker-compose stop
docker-compose rm -f

# remove a image
docker rmi ibge_cassandra:1.0 -f

# Constroi a imagem docker
docker build -t ibge_cassandra:1.0 .

# Executa os containers conforme docker-compose.yml
docker-compose up -d

