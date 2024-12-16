# Тестовое задание для "Medods"

## Запуск приложения

### 1. Переменные окружения
Необходимо создать .env файл в корне проекта с переменными
```
ENV=local
POSTGRES_DB=medods
POSTGRES_USER=postgres
POSTGRES_PASSWORD=password
JWT_SIGN=jwtsign
JWT_EXP=15m
DATABASE_URL=postgres://postgres:password@postgres:5432/medods?sslmode=disable
DATABASE_URL_MIGRATIONS=postgres://postgres:password@localhost:5432/medods?sslmode=disable

```

### 2. Запуск в докере
Приложение запускается командой `make`
