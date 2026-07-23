# Итоговый проект по курсу Golang-разработчик

Простое веб-приложение для управления списком задач, позволяющее пользователям создавать,
редактировать, отмечать выполненными и удалять планируемые регулярные и одноразовые задачи. 

## Описание содержания

- Директория `dbase` содержит базу данных scheduler.db 

- В директории `tests` находятся тесты для проверки API, которое должно быть реализовано в веб-сервере.

- Директория `web` содержит файлы фронтенда.

- Директория `pkg/api` содержит хэндлеры, их регистрацию и сопутствующие функции

- Директория `pkg/db` содержит функции требуемые для взаимодействием с базой данных

## Выполнено задание повышенной сложности выбора каталога для базы данных и порта сервера в переменных окружения в файле .env

## Конфигурация файла settings.go для проекта
```
package tests
var Port = 7540
var DBFile = "../dbase/scheduler.db"
var FullNextDate = false
var Search = false
var Token = ``
```

## Список пройденных тестов
```
go test -run ^TestNextDate$ ./tests
go test -run ^TestAddTask$ ./tests
go test -run ^TestTasks$ ./tests
go test -run ^TestTask$ ./tests
go test -run ^TestEditTask$ ./tests
go test -run ^TestDone$ ./tests
go test -run ^TestDelTask$ ./tests
go test ./tests
```
## Запуск
Запуск производится посредством команды 
```go run main.go```