# Para executar o container do mssql
docker-compose up

# Cria o banco de dados dbpaises e a tabela paises
cat paises.sql | docker exec -i api_dfe_paises_sqlserver_1 bash -c '/opt/mssql-tools/bin/sqlcmd -U sa -P Senha123'

# Start SQL Server connector
curl -i -X POST -H "Accept:application/json" -H  "Content-Type:application/json" http://localhost:8083/connectors/ -d @register-sqlserver.json

# Consume messages from a Debezium topic
docker-compose -f docker-compose.yaml exec kafka /kafka/bin/kafka-console-consumer.sh \
    --bootstrap-server kafka:9092 \
    --from-beginning \
    --property print.key=true \
    --topic server1.dbo.paises

# Modify records in the database via SQL Server client (do not forget to add `GO` command to execute the statement)
docker-compose -f docker-compose.yaml exec sqlserver bash -c '/opt/mssql-tools/bin/sqlcmd -U sa -P senha123 -d dbpaises'

# Shut down the cluster
docker-compose -f docker-compose.yaml down


