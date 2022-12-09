# Golang RESTful Web API with MongoDB  

Este exemplo mostra uma API Restfull CRUD, utilizando como banco de dados o MongoDB.

O projeto foi estruturado seguindo Arquitetura Limpa, demonstrando o desenvolvimento de uma API Restfull utilizando dois framework web,
o Echo e o Fiber. As duas implementações utilizam o mesmo dominio e também compartilham o acesso ao banco de dados, tudo graças a aplicação da Arquitetura Limpa.
O que muda realmente é a parte de handlers e a configuração de rotas e servidor.

Também foi incluíndo o swagger para documentação da API, preferi baixar o swagger-ui e adicionar o JSON que descreve a API, em vez de usar as soluções que geram
automaticamente o swagger a partir de comentários no código fonte.

Também foi adicionado suporte a observalidade utilizando OpenTelemetry.

### Este exemplo utiliza as seguintes tecnologias:
 - GO versão 1.18
 - Echo web framework 
 - Fiber web framework
 - MongoDB última versão
 - OpenApi v3
 - OpenTelemetry
 - Docker
 - Docker compose

### Para executar a aplicação faça os seguintes passos:

#### O primeiro passo e fazer um clone do repositório, e na raiz do projeto executar um dos comandos abaixo. 
#### Observação: É necessário ter o docker e docker-compose instalado no seu computador.

 - Para executar a versão utilizando o framework Echo use o seguinte comando:
```bash
 docker-compose -f deployments/docker-compose-echo-local.yaml up -d --build
```
E para finalizar você pode utilizar este comando
```bash
 docker-compose -f deployments/docker-compose-echo-local.yaml down
```

ou

 - Para executar a versão utilizando o framework Fiber use o seguinte comando:
```bash
 docker-compose -f deployments/docker-compose-fiber-local.yaml up -d --build
```
E para finalizar você pode utilizar este comando
```bash
 docker-compose -f deployments/docker-compose-fiber-local.yaml down
```


