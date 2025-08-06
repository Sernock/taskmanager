# TaskManager
Task manager API на Go

Простой CRUD-проект на Go с использованием Gin и SQLite.  
Позволяет создавать, просматривать, обновлять и удалять задачи.

## Структура проекта

C:.
│   .gitignore
│   go.mod
│   go.sum
│   main.go
│   README.md
│   taskmanager.db
│
└───internal
    ├───db
    │       db.go
    │       tasks.go
    │
    ├───handlers
    │       task.go
    │
    └───models
            task.go


- `main.go` — точка входа, настройка маршрутов.
- `internal/models` — структуры данных (`Tasks`).
- `internal/db` — функции работы с базой данных.
- `internal/handlers` — HTTP-обработчики.
- `internal/db/tasks.db` — база данных SQLite (создаётся автоматически при первом запуске).

## Запуск проекта

1. Клонировать репозиторий:

git clone https://github.com/username/taskmanager.git
cd taskmanager

2. Установить зависимости: 
go mod tidy

3. Запустить сервер:

go run main.go

Сервер будет слушать http://localhost:8080.

- Проект использует SQLite для хранения данных — база `taskmanager.db` создаётся автоматически.
- Gin работает в режиме `debug` по умолчанию. Для production рекомендуется выставить `GIN_MODE=release`.
- Все запросы и ответы возвращаются в формате JSON.
