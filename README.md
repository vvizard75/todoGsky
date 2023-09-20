1. Получить все задачи пользователя
Метод: GET

URL: /users/{userId}

Параметры:

userId - идентификатор пользователя.

Пример запроса:
curl -v -X GET http://localhost:8080/users/1

Пример ответа:
HTTP/1.1 200 OK
[
    {
        "Id": 1,
        "Name": "Task 1",
        "Description": "Description for task 1",
        "Completed": false
    },
    {
        "Id": 2,
        "Name": "Task 2",
        "Description": "Description for task 2",
        "Completed": false
    }
]



2. Создать новую задачу
Метод: POST

URL: /users/{userId}

Параметры:

userId - идентификатор пользователя.

Тело запроса:
{
    "Name": "Task Name",
    "Description": "Task Description"
}
Пример запроса:
curl -v -X POST -H "Content-Type: application/json" -d '{"Name": "New Task", "Description": "Description for the new task"}' http://localhost:8080/users/1

Пример ответа:
HTTP/1.1 201 Created
Id

3. Обновить существующую задачу
Метод: PUT

URL: /users/{userId}/{taskId}

Параметры:

userId - идентификатор пользователя.
taskId - идентификатор задачи.
Тело запроса:
{
    "Name": "Updated Task Name",
    "Description": "Updated Task Description"
}
Пример запроса:
curl -v -X PUT -H "Content-Type: application/json" -d '{"Name": "Updated Task", "Description": "Updated description for the task"}' http://localhost:8080/users/1/3

Пример ответа:
HTTP/1.1 200 OK


4. Удалить задачу
Метод: DELETE

URL: /users/{userId}/{taskId}

Параметры:

userId - идентификатор пользователя.
taskId - идентификатор задачи.
Пример запроса:
curl -v -X DELETE http://localhost:8080/users/1/3

Пример ответа:
HTTP/1.1 200 OK

