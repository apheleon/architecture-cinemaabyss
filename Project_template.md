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
[скриншот]([URL](https://disk.yandex.ru/i/INmyWdv2e3Llpg), "результат npm run test:kubernetes")
[скриншот]([URL](https://disk.yandex.ru/i/OZzO9lr3MX22sA) "ответ https://cinemaabyss.example.com/api/movies")
[скриншот]([URL](https://disk.yandex.ru/i/XKka1gqzZCD1eA) "лог event-service")


# Задание 4

Правки в helm внесены