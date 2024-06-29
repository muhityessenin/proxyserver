# Proxy Server - это простой прокси сервер который принимает запрос в формате JSON, взаимодействует с внешними сервисами и возвращает ответ пользователю в том же формате.


## Установка и запуск

### Требования

- Docker
- Docker Compose

### Запуск сервера

1. Склонируйте репозиторий:

  git clone https://github.com/muhityessenin/proxyserver.git

2. Постройте и запустите контейнеры:

  make build

  make up

  Сервер будет доступен по адресу http://localhost:8080/proxy

3. Отправьте запрос в формате JSON, и в ответ вы получаете в таком же формате со всеми заголовками HTTP. 

