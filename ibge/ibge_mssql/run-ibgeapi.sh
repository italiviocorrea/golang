#!/usr/bin/env bash

# Constroi a ibgeapims
go build

# Para e remove os container em execução
docker-compose stop
docker-compose rm -f

# remove a image
docker rmi ibge_mssql:1.1 -f

# Constroi a imagem docker
docker build -t ibge_mssql:1.1 .

# Executa os containers conforme docker-compose.yml
docker-compose up -d

