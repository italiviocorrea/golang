FROM icorrea/alpine:3.5-pt-br
MAINTAINER Italivio Correa <italiviocorrea@gmail.com>

# variaveis de ambiente com valores default
ENV CASSANDRA_API_SERVER=0.0.0.0:8080 \
    CASSANDRA_API_DBHOST="localhost:19042,localhost:29042,localhost:39042" \
    CASSANDRA_API_DBUSER=admin \
    CASSANDRA_API_DBPWD=senha#123 \
    CASSANDRA_API_DATABASE=ibgeapi \
    CASSANDRA_API_CONTEXT=/ibge/v3 \
    CASSANDRA_API_LOGLEVEL=4

# adiciona a API
ADD ibge_cassandra /ibge_cassandra

# ADD security /security

# adiciona o site do swagger-ui
ADD swagger-ui /swagger-ui

# define a porta exposta pela API
EXPOSE 8080

ENTRYPOINT ["./ibge_cassandra"]