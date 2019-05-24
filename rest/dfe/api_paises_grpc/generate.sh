#!/bin/bash

protoc ./paisespb/paises.proto --go_out=plugins=grpc:.
protoc -I ./api_paises_grpc/ api_paises_grpc/paisespb/paises.proto  --go_out=plugins=grpc:./api_paises