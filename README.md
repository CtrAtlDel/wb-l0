## WB project (L0)

> Как запустить проект (how)

Чтобы запустить проект необходимо:
1. Поднять базу данных в докере, для этого выполняем из корня проекта:
```bash
cd docker/
docker-compose up -d
```
2. Установить с помощью пакетного менеджера (для примера взят brew) nats-streaming-server:
```bash
brew install nats-streaming-server
```
3. Необходимо запустить сервер выполнив команду:
```bash 
nats-streaming-server
```
4. Собрать и запустить основное приложение:
```bash
go build ivankvasov/project/cmd/main                                                       
go run ivankvasov/project/cmd/main
```
5. Собрать и запустить паблишер:
```bash
go build ivankvasov/publisher/cmd
go run ivankvasov/publisher/cmd
```
5. *Или с помощью makefile:
```bash
make
make publisher
```
В базе данные хранятся в одной таблице в формате JSON и извлекаются по ключу order_uid