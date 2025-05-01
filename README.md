# TODO List API на Go + Fiber

## 📌 Описание проекта

REST API-сервис для управления задачами (TODO-лист), реализованный на Go с использованием фреймворка Fiber и базой данных PostgreSQL.

Реализованные возможности:
- ✅ Создание задачи (`POST /tasks`)
- ✅ Получение списка задач (`GET /tasks`)
- ✅ Обновление задачи (`PUT /tasks/:id`)
- ✅ Удаление задачи (`DELETE /tasks/:id`)

## 🧱 Используемый стек

- Go + Fiber
- PostgreSQL (через `pgx`)
- Docker (опционально)

## 🗃 Структура базы данных

Создаётся таблица `tasks` со следующими полями:

| Поле         | Тип      | Описание                                           |
|--------------|----------|----------------------------------------------------|
| id           | SERIAL   | Первичный ключ                                     |
| title        | TEXT     | Название задачи (обязательно)                      |
| description  | TEXT     | Описание задачи                                    |
| status       | TEXT     | Статус (`new`, `in_progress`, `done`) по умолчанию `new` |
| created_at   | TIMESTAMP| Дата создания (по умолчанию `now()`)              |
| updated_at   | TIMESTAMP| Дата обновления (по умолчанию `now()`)            |

---

## 🚀 Запуск проекта локально (без Docker)

1. Клонируйте репозиторий:

```bash
git clone <your-repo-url>
cd <project-directory>
```

2. Создайте `.env` файл в корне проекта со следующим содержимым:

```env
DATABASE_URL="postgres://user:password@localhost:5432/todo-list"
```

3. Запустите приложение:

```bash
go run main.go
```

4. Перейдите в браузере или через Postman:

```
http://localhost:3000/tasks
```

---

## ⚙️ Выполнение миграций

1. Добавьте дополнительные переменные окружения в `.env`:

```env
PG_URL="postgres://user:password@localhost:5432"
DB_NAME="todo-list-migrated"
```

2. Запустите приложение с флагом миграции:

```bash
go run main.go --migrate
```

---

## 📦 Запуск через Docker

1. Убедитесь, что у вас установлен Docker и Docker Compose.

2. Добавьте следующие переменные окружения в `.env` файл:

```env
POSTGRES_USER=user
POSTGRES_PASSWORD=password
POSTGRES_DB=docker-todo-list
DATABASE_URL=postgres://user:password@db:5432/todo-list
PG_URL=postgres://user:password@db:5432
DB_NAME=docker-todo-list
```

💡 **Важно:** в переменных окружения для Docker необходимо использовать `db` вместо `localhost` в строках подключения (`DATABASE_URL`, `PG_URL`), так как `db` — это имя сервиса базы данных в `docker-compose.yml`.

3. Соберите и запустите контейнеры:

```bash
docker compose up --build
```

Приложение будет доступно по адресу: [http://localhost:3000/tasks](http://localhost:3000/tasks)

---

## 🖼️ Скриншоты работы API

### ✅ Создание задачи (POST)
![Создание задачи](docs/Отправка POST с записью в бд.png)

### 📋 Список задач (GET)
![Список задач](docs/Вывод всех задач.png)

### ✏️ Обновление задачи (PUT)
![Обновление задачи](docs/Изменение статуса задачи запросом PUT.png)

![Вывод обновленной задачи](docs/docs/Вывод измененной задачи.png)

### ❌ Удаление задачи (DELETE)
![Удаление задачи](docs/Запуск delete запроса.png)

![Список с удаленной задачей](docs/Вывод списка задач с удаленной задачей.png)

### 🧾 Подтверждение в базе данных (pgAdmin)
![pgAdmin](docs/Вывод задач в pgAdmin.png)

### ⚠️ Ошибка при неверном POST
![Ошибка POST](docs/Проверка ошибки с неправильным содержанием POST запроса.png)

---
