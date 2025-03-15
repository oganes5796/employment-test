# 📝 Task Manager API

Task Manager API — это REST API-сервис для управления задачами (TODO-лист) на основе Go, Fiber и PostgreSQL. Проект следует принципам чистой архитектуры и стандартам Uber Go Style Guide.

---

## 🚀 Функциональность
✅ Создание задачи  
✅ Получение списка всех задач  
✅ Обновление задачи  
✅ Удаление задачи  
✅ Полноценная Swagger-документация  
✅ Логирование с Zap  
✅ Подключение к PostgreSQL через pgx  
✅ Миграции через Goose  
✅ Тесты с покрытием  

---

## 🛠️ Стек технологий
- Go (Fiber)
- PostgreSQL (pgx)
- Goose (миграции)
- Zap (логирование)
- Swagger (документация)
- Testify (тестирование)

---

## 📦 Установка и запуск
### 1. Клонирование репозитория
```bash
git clone <URL-РЕПОЗИТОРИЯ>
cd task-manager
```

### 2. Установка зависимостей
Установите все необходимые зависимости командой:
```bash
make install-deps
```

### 3. Создание `.env` файла
Создайте `.env` файл в корне проекта со следующим содержимым:
```
POSTGRES_DB=user-service
POSTGRES_USER=user
POSTGRES_PASSWORD=password
MIGRATION_DIR=./migrations
PG_DSN="host=localhost port=5435 dbname=user-service user=user password=password sslmode=disable"
LOCAL_HOST=localhost
LOCAL_PORT=8080
```

### 4. Запуск базы данных
```bash
docker-compose up -d
```

### 6. Запуск программы
```bash
go run cmd/main.go
```

---

## 📚 Документация API (Swagger)
После запуска перейдите по адресу:
```
http://localhost:8080/swagger/index.html
```

---

## 🔍 Примеры работы в Postman
### ➕ Создание задачи
**POST** `/tasks`  
**Body (JSON):**
```json
{
    "title": "Купить продукты",
    "description": "Купить молоко, яйца и хлеб",
    "status": "new"
}
```

### 📋 Получение всех задач
**GET** `/tasks`  

### ✏️ Обновление задачи
**PUT** `/tasks/{id}`  
**Body (JSON):**
```json
{
    "title": "Обновленная задача",
    "description": "Описание обновлено",
    "status": "in_progress"
}
```

### ❌ Удаление задачи
**DELETE** `/tasks/{id}`  

---

## 🧪 Запуск тестов
```bash
go test ./... -cover
```

Для просмотра покрытия кода тестами:
```bash
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out
```

---

## 🤝 Контакты
Если у вас возникли вопросы или предложения по улучшению проекта, пишите в Issues или создавайте Pull Request.

---

## 📜 Лицензия
Проект распространяется по лицензии MIT.

