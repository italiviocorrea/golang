## Orientações para publicação no Minikube

Para testar a aplicação em ambiente kubernetes, temos os arquivos para implantação do mongodb, mongo-express e também a aplicação, nas
duas versões Echo e Fiber.

### Pré requisitos para publicar a aplicação no minikube

É necessário fazer a instalação do minikube conforme a SO que você está utilizando. Conforme o link abaixo
https://minikube.sigs.k8s.io/docs/start/

### Passos para publicar

- Iniciar o Minikube
```shell
minikube start
```
- Instalar o mongodb no minikube:
    - Utilizando o kubectl vamos instalar as configurações de volumes, segredos, configurações, etc.
```shell
        kubectl apply -f deployments/k8s/mongodb-pv.yaml
        kubectl apply -f deployments/k8s/mongodb-pvc.yaml
        kubectl apply -f deployments/k8s/mongodb-secret.yaml
        kubectl apply -f deployments/k8s/mongodb-configmap.yaml
        kubectl apply -f deployments/k8s/mongodb-deployment.yaml
        kubectl apply -f deployments/k8s/mongo-express-deployment.yaml
        kubectl get all | grep mongo
```
- Vamos criar a imagem do Docker no minikube
  - Definir o ambiente, para que o daemon Docker execute na instância do Minikube
```shell
    eval $(minikube docker-env)
```
- E finalmente criar a imagem Docker no Minikube
```shell
     docker build -f build/deploy/Dockerfile.fiber apisdt-fiber-go:local .
```
- Com a imagem criada, vamos implantar a aplicação no minikube
```shell
     kubectl apply -f deployments/k8s/apisdt-fiber-local.yaml
```
- Agora vamos liberar a porta 8080 para o localhost

```shell
  kubectl port-forward svc/apisdt-echo-go-service 8080:8080
```

- Utilize o seu browse favorito para acessar a aplicação

```http request
  http://localhost:8080/api/v1/swagger/index.html
```

Pronto agora temos uma aplicação completa com banco de dados persistente implantanda no kubernetes.
