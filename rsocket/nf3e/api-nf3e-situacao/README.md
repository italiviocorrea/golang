# 💻 API NF3e consulta situação

Consulta a NF3e através de uma chave informada, no banco de dados do YugabyteDB utilizando a API CQL.
O protocolo de comunicação utilizado é rsocket.

### ⚙️ Descrição do serviço

- **Serviço:** api-nf3e-consulta-situacao
- **Função:** API destinado a devolver os dados de protocolo e eventos da NF3e.
- **Protocolo:** RSocket com fluxo do tipo **Request-Response**.
- **Método:** nenhum
- **Parâmetro da Mensagem :**
    - **Entrada :**
        - **chNF3e:** Chave de acesso da NF3e a consultada a situação
    - **Retorno :**
        - **NF3eSituacao:**
            - **chNF3e:** Chave de acesso da NF3e
            - **versao:** Versão da NF3e
            - **tpAmb:** Tipo de ambiente.
            - **cStat:** Codigo do status da NF3e.
            - **xMotivo:** Motivo da NF3e
            - **cUF:** Codigo da UF que autorizou a NF3e
            - **protNF3e:** String XML do protocolo da autorização da NF3e.
            - **procEventoNF3e:** Lista de XML de eventos autorizados e relacionados a NF3e (opcional)

## 🛠 Tecnologias

As seguintes ferramentas foram utilizadas na construção desta API:

- **[Go](https://golang.org/)**
- **[Go Rsocket](https://github.com/rsocket/rsocket-go)** 
- **[Go Cql](https://github.com/yugabyte/gocql)** 
- **[RSocket Client CLI (RSC)](https://github.com/making/rsc)** Ferramenta cliente para testar API rsocket via linha de comando, semelhante ao comando curl.


## 🚀 Exemplos de como consumir a API, utilizando o utilitário RSC:

💡 Para executar os exemplos de 1 a 4 é necessário o java e o utilitário rsc-<version>.jar .

1. Exemplo de uma NF3e não existente na base de dados:

   1.1. Comando: ```java -jar /home/icorrea/tools/rsocket-client/rsc-0.6.1.jar tcp://localhost:7878 -d '50210602935843000105660016292451931647934135'```

   1.2. Resposta : ``` not found```

2. Exemplo de uma NF3e existente na base de dados:

   2.1. Comando: ```java -jar /home/icorrea/tools/rsocket-client/rsc-0.6.1.jar tcp://localhost:7878 -d '50210602935843000105660016292451931647934134'```

   2.2. Resposta : ```{"chNF3e":"50210602935843000105660016292451931647934134","versao":"1.00","tpAmb":"2","cStat":"100","xMotivo":"Autorizado o Uso da NF3e","cUF":"50","protNF3e":"<protNF3e versao=\"1.00\" xmlns=\"http://www.portalfiscal.inf.br/nf3e\" ><infProt><tpAmb>2</tpAmb><verAplic>0.0.28</verAplic><chNF3e>50210602935843000105660016292451931647934134</chNF3e><nProt>150210000000527</nProt><dhRecbto>2021-06-30 09:36:12-04:00</dhRecbto><digVal>cmhpMzdBM0FpTUYvOTQ5R0F1WjJhU2xiKzVBPQ==</digVal><cStat>100</cStat><xMotivo>Autorizado o Uso da NF3e</xMotivo></infProt></protNF3e>"}```




