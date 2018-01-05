FROM icorrea/alpine:3.5-pt-br
MAINTAINER Italivio Correa <italiviocorrea@gmail.com>

# variaveis de ambiente com valores default
ENV MONGODB_API_SERVER="0.0.0.0:8080" \
    MONGODB_API_MONGODBHOST="localhost" \
    MONGODB_API_DBUSER="" \
    MONGODB_API_DBPWD="" \
    MONGODB_API_DATABASE="ibgeapi" \
    MONGODB_API_CONTEXT="/ibge/v3" \
    MONGODB_API_LOGLEVEL=4

# adiciona a API
ADD ibge_mongodb /ibge_mongodb

# adiciona o site do swagger-ui
ADD swagger-ui /swagger-ui

# define a porta exposta pela API
EXPOSE 8080

ENTRYPOINT ["./ibge_mongodb"]
