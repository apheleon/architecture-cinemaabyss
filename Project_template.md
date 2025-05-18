## Изучите [README.md](.\README.md) файл и структуру проекта.

# Задание 1

Диаграмма внутри проекта: src\diagrams\c4_containers.puml

# Задание 2

### 1. Proxy
Proxy сервис реализован, на Go, само собой. Папка с сервисом: src\microservices\proxy

### 2. Kafka
[Скриншоты]([URL](https://disk.yandex.ru/d/3DLC87xSC1DpjA) "папка со скриншотами, яндекс диск")
 
# Задание 3

Билд прошёл успешно, [скриншот]([URL](https://disk.yandex.ru/i/OvODIZk7KEGc4g), "build-and-run success")

### Proxy в Kubernetes

#### Шаг 1

  Конфиг в проекте .docker\config.json

#### Шаг 2

[скриншот]([URL](https://disk.yandex.ru/i/dXaixZji8vp_6g) "все поды подняты")
 
  
#### Шаг 3
[скриншот]([URL](https://disk.yandex.ru/i/5Ik5OzYcC1pB4Q), "результат npm run test:kubernetes")
[скриншот]([URL](https://disk.yandex.ru/i/OZzO9lr3MX22sA) "ответ https://cinemaabyss.example.com/api/movies")
[скриншот]([URL](https://disk.yandex.ru/i/XKka1gqzZCD1eA) "лог event-service")


# Задание 4
Для простоты дальнейшего обновления и развертывания вам как архитектуру необходимо так же реализовать helm-чарты для прокси-сервиса и проверить работу 

Для этого:
1. Перейдите в директорию helm и отредактируйте файл values.yaml

```yaml
# Proxy service configuration
proxyService:
  enabled: true
  image:
    repository: ghcr.io/db-exp/cinemaabysstest/proxy-service
    tag: latest
    pullPolicy: Always
  replicas: 1
  resources:
    limits:
      cpu: 300m
      memory: 256Mi
    requests:
      cpu: 100m
      memory: 128Mi
  service:
    port: 80
    targetPort: 8000
    type: ClusterIP
```

- Вместо ghcr.io/db-exp/cinemaabysstest/proxy-service напишите свой путь до образа для всех сервисов
- для imagePullSecret проставьте свое значение (скопируйте из конфигурации kubernetes)
  ```yaml
  imagePullSecrets:
      dockerconfigjson: ewoJImF1dGhzIjogewoJCSJnaGNyLmlvIjogewoJCQkiYXV0aCI6ICJaR0l0Wlhod09tZG9jRjl2UTJocVZIa3dhMWhKVDIxWmFVZHJOV2hRUW10aFVXbFZSbTVaTjJRMFNYUjRZMWM9IgoJCX0KCX0sCgkiY3JlZHNTdG9yZSI6ICJkZXNrdG9wIiwKCSJjdXJyZW50Q29udGV4dCI6ICJkZXNrdG9wLWxpbnV4IiwKCSJwbHVnaW5zIjogewoJCSIteC1jbGktaGludHMiOiB7CgkJCSJlbmFibGVkIjogInRydWUiCgkJfQoJfSwKCSJmZWF0dXJlcyI6IHsKCQkiaG9va3MiOiAidHJ1ZSIKCX0KfQ==
  ```

2. В папке ./templates/services заполните шаблоны для proxy-service.yaml и events-service.yaml (опирайтесь на свою kubernetes конфигурацию - смысл helm'а сделать шаблоны для быстрого обновления и установки)

```yaml
template:
    metadata:
      labels:
        app: proxy-service
    spec:
      containers:
       Тут ваша конфигурация
```

3. Проверьте установку
Сначала удалим установку руками

```bash
kubectl delete all --all -n cinemaabyss
kubectl delete  namespace cinemaabyss
```
Запустите 
```bash
helm install cinemaabyss .\src\kubernetes\helm --namespace cinemaabyss --create-namespace
```
Если в процессе будет ошибка
```code
[2025-04-08 21:43:38,780] ERROR Fatal error during KafkaServer startup. Prepare to shutdown (kafka.server.KafkaServer)
kafka.common.InconsistentClusterIdException: The Cluster ID OkOjGPrdRimp8nkFohYkCw doesn't match stored clusterId Some(sbkcoiSiQV2h_mQpwy05zQ) in meta.properties. The broker is trying to join the wrong cluster. Configured zookeeper.connect may be wrong.
```

Проверьте развертывание:
```bash
kubectl get pods -n cinemaabyss
minikube tunnel
```

Потом вызовите 
https://cinemaabyss.example.com/api/movies


## Удаляем все

Установите https://istio.io/latest/docs/reference/commands/istioctl/

```bash
kubectl delete all --all -n cinemaabyss
kubectl delete namespace cinemaabyss
```