# GO-project

ЛР2: сервер с подключением к PostgreSQL, инициализацией таблиц и тестовым POST-хэндлером `/dbtest`.

## Запуск БД

```bash
docker compose up -d
```

По умолчанию сервер подключается к:

```text
postgres://postgres:postgres@localhost:5432/mini_avito?sslmode=disable
```

Можно переопределить строку подключения переменной окружения:

```bash
export POSTGRES_DSN="postgres://postgres:postgres@localhost:5432/mini_avito?sslmode=disable"
```

## Запуск сервера

```bash
go mod tidy
go run .
```

## Проверка

```bash
curl -X POST http://localhost:8080/dbtest -d 'hello database'
```

Ожидаемый ответ:

```json
{
  "id": 1,
  "body": "hello database",
  "created_at": "..."
}
```

## Таблицы

При старте создаются таблицы:

- `users` — пользователи мини-Avito;
- `ads` — объявления со статусами `active` / `sold`;
- `db_test` — тестовая таблица для проверки записи тела запроса.
