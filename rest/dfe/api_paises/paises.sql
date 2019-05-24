-- Cria o banco de dados dbpaises
CREATE DATABASE dbpaises;
GO
USE dbpaises;
-- Ativa o CDC para o banco de dados
EXEC sys.sp_cdc_enable_db;

-- Create a tabela paises
CREATE TABLE paises ( 
  codigo INT NOT NULL PRIMARY KEY, 
  nome VARCHAR(32) NOT NULL, 
  inicio_vigencia DATETIME NOT NULL, 
  fim_vigencia DATETIME
);
-- Ativa o CDC para a tabela Paises
EXEC sys.sp_cdc_enable_table @source_schema = 'dbo', @source_name = 'paises', @role_name = NULL, @supports_net_changes = 0;
GO
