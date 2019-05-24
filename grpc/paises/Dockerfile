FROM icorrea/alpine:3.9
MAINTAINER Italivio Correa <italiviocorrea@gmail.com>

# variaveis de ambiente com valores default
ENV API_SRV_ADDR=0.0.0.0:8080 \
    API_DB_HOST=dbpaises \
    API_DB_USER=sa \
    API_DB_PWD=Senha123 \
    API_DB_PORT=1433 \
    API_DB_NAME=dbpaises \
    API_SRV_CONTEXT=/dfe/v1 \
    API_SRV_LOGLEVEL=4

# adiciona a API
ADD main /api_paises

# ADD security /security

# adiciona o site do swagger-ui
ADD swagger /swagger

# define a porta exposta pela API
EXPOSE 8080

ENTRYPOINT ["./api_paises"]
