# TaskManager

**TaskManager** — простой REST API для управления задачами, написанный на Go с использованием фреймворка [Gin](https://github.com/gin-gonic/gin) и JWT-аутентификации.

---

## Описание

Этот проект позволяет создавать, читать, обновлять и удалять задачи (CRUD) с защитой маршрутов через JWT токены. Идеально подходит для обучения или как стартовая база для более сложных приложений.

---

## Возможности

- Регистрация и логин пользователей с JWT аутентификацией
- Создание, получение, обновление и удаление задач
- Авторизация защищенных маршрутов через JWT Middleware
- Хранение данных в SQLite (или любой другой поддержуемой базе данных)

---

## Технологии

- Go 1.20+
- Gin — веб-фреймворк
- JWT ([github.com/golang-jwt/jwt/v5](https://github.com/golang-jwt/jwt))
- SQLite (или другая СУБД)
- godotenv для работы с переменными окружения

---

## Установка и запуск

1. Клонируйте репозиторий:

-  git clone https://github.com/yourusername/taskmanager.git
-  cd taskmanager

2. Создайте файл .env в корне проекта и добавьте в него:

- JWT_SECRET=your_secret_key

3. Установите зависимости:

- go mod tidy

4. Запустите программу:

- go run main.go

5. Api будет доступен по адресу:

- API будет доступен по адресу: http://localhost:8080

---

## Использование API

1. Регистрация:

POST /register
Content-Type: application/json

{
  "username": "user1",
  "password": "password123"
}

2. Логин:

POST /login
Content-Type: application/json

{
  "username": "user1",
  "password": "password123"
}

Ответ содержит токен:

Формат типа:

{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}

3. Работа с задачами(требует авторизацию):

<small>Добавьте в заголовок запроса:<small>

- Authorization: Bearer <ваш_jwt_токен>

Получить задачи:

Get /tasks

Создать задачу:

POST /tasks
Content-Type: application/json

{
  "title": "Новая задача",
  "description": "Описание задачи",
  "completed": false
}


Обновить задачу:

PUT /tasks/:id
Content-Type: application/json

{
  "title": "Обновленная задача",
  "description": "Новое описание",
  "completed": true
}

Удалить задачу:

DELETE /tasks/:id


---

## Структура проекта

taskmanager/
├── internal/
│   ├── handlers/      # HTTP-обработчики
│   ├── middleware/    # Middleware для аутентификации
│   ├── db/            # Работа с базой данных
│   └── models/        # Модели данных
├── .env               # Переменные окружения
├── main.go            # Точка входа в приложение
├── go.mod
└── README.md



Не храните секреты в репозитории — используйте .env

Для продакшена настройте HTTPS и надежное хранение ключей

Добавьте валидацию входящих данных для безопасности