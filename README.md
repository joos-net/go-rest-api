# Создание REST API‑сервиса (CRUD задач)
## Островский Евгений

REST API для управления списком задач с хранением в памяти.

## Возможности

- Создание, чтение, обновление и удаление задач
- JSON-ответы с HTTP-статусами
- Логирование запросов
- Эндпоинт для проверки состояния сервиса

## Эндпоинты API

| Метод | Эндпоинт | Описание | Коды статусов |
|-------|----------|----------|---------------|
| GET | /tasks | Получить все задачи | 200 |
| POST | /tasks | Создать новую задачу | 201, 400 |
| GET | /tasks/{id} | Получить задачу по ID | 200, 404 |
| PUT | /tasks/{id} | Обновить задачу по ID | 200, 400, 404 |
| DELETE | /tasks/{id} | Удалить задачу по ID | 204, 404 |
| GET | /health | Проверка состояния сервиса | 200 |

## Модель задачи

```json
{
  "id": 1,
  "title": "Название задачи",
  "done": false,
  "created_at": "2024-01-01T12:00:00Z"
}

## Примеры запросов
POST
```
curl -X POST http://localhost:8080/tasks \ 
  -H "Content-Type: application/json" \
  -d '{"title":"Create REST API","done":false}'
```
![1](https://github.com/joos-net/go-rest-api/blob/main/1.png)

GET
```
curl http://localhost:8080/tasks
```
![2](https://github.com/joos-net/go-rest-api/blob/main/2.png)

PUT
```
curl -X PUT http://localhost:8080/tasks/1 \ 
  -H "Content-Type: application/json" \
  -d '{"title":"Создание REST API","done":true}'
```
![3](https://github.com/joos-net/go-rest-api/blob/main/3.png)

DELETE
```
curl -X DELETE http://localhost:8080/tasks/1
curl http://localhost:8080/tasks/1

{"error":"Задача с ID 1 не найдена"}
```
![4](https://github.com/joos-net/go-rest-api/blob/main/4.png)

GET health
```
curl http://localhost:8080/health          
{"status":"healthy"}
```
![5](https://github.com/joos-net/go-rest-api/blob/main/5.png)

![6](https://github.com/joos-net/go-rest-api/blob/main/6.png)