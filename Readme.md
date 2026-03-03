# Organizational Structure API  
# API организационной структуры

REST API для управления иерархией подразделений и сотрудников.

---

## 🚀 Запуск / Run

### 🐳 Через Docker (рекомендуется / Recommended)

```bash
docker compose up --build
```

Приложение будет доступно по адресу:

http://localhost:8085

---

## 🗄 Локальный запуск (без Docker)

Необходимо создать базу данных:

test_for_work

### Запуск миграций

Установить goose:

```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
```

Применить миграции:

```bash
goose -dir migrations postgres "host=localhost user=postgres password=YOUR_PASSWORD dbname=test_for_work port=5432 sslmode=disable" up
```

---

# 📦 Примеры JSON-запросов

## ➕ Создание отдела

POST /departments

```json
{
  "name": "Dwayne",
  "parent_id": null
}
```

---

## ➕ Создание сотрудника

POST /departments/{id}/employees

```json
{
  "full_name": "Vin Diesel",
  "position": "Action",
  "hired_at": "1980-01-01"
}
```

---

# 🚀 Departments API

## ➕ Создать отдел

```bash
curl -X POST http://localhost:8085/departments -H "Content-Type: application/json" -d "{\"name\":\"Backend\"}"
```

## 📄 Получить список отделов

```bash
curl http://localhost:8085/departments
```

## 🔍 Получить отдел по ID

```bash
curl http://localhost:8085/departments/1
```

## ✏ Обновить отдел

```bash
curl -X PUT http://localhost:8085/departments/1 -H "Content-Type: application/json" -d "{\"name\":\"Backend Updated\"}"
```

## ❌ Удалить отдел

```bash
curl -X DELETE http://localhost:8085/departments/1
```

---

# 👨‍💻 Employees API

## ➕ Создать сотрудника в отделе

```bash
curl -X POST http://localhost:8085/departments/1/employees -H "Content-Type: application/json" -d "{\"full_name\":\"Ivan Ivanov\",\"position\":\"Golang Developer\",\"hired_at\":\"2020-01-01\"}"
```

## 📄 Получить всех сотрудников отдела

```bash
curl http://localhost:8085/departments/1/employees
```

## 🔍 Получить сотрудника по ID

```bash
curl http://localhost:8085/employees/1
```

## ✏ Обновить сотрудника

```bash
curl -X PUT http://localhost:8085/employees/1 -H "Content-Type: application/json" -d "{\"full_name\":\"Ivan Petrov\",\"position\":\"Senior Golang Developer\",\"hired_at\":\"2020-01-01\"}"
```

## ❌ Удалить сотрудника

```bash
curl -X DELETE http://localhost:8085/employees/1
```