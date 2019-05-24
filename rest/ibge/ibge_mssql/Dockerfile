FROM icorrea/alpine:3.5-pt-br
MAINTAINER Italivio Correa <italiviocorrea@gmail.com>

# variaveis de ambiente com valores default
ENV MSSQL_API_SERVER=0.0.0.0:8080 \
    MSSQL_API_DBHOST=dbibgeapims \
    MSSQL_API_DBUSER=sa \
    MSSQL_API_DBPWD=senha#123 \
    MSSQL_API_DBPort=1433 \
    MSSQL_API_DATABASE=dbibgeapi \
    MSSQL_API_CONTEXT=/ibge/v3 \
    MSSQL_API_LOGLEVEL=4

# adiciona a API
ADD ibge_mssql /ibge_mssql

# ADD security /security

# adiciona o site do swagger-ui
ADD swagger-ui /swagger-ui

# define a porta exposta pela API
EXPOSE 8080

ENTRYPOINT ["./ibge_mssql"]