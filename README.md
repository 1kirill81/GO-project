# GO-project — ЛР2 Docker

ЛР2 сделана поверх ЛР1 и теперь полностью запускается через Docker Compose: и Go-сервер, и PostgreSQL.

## Что осталось от ЛР1

- `GET /test` проходит через 3 слоя `handler -> service -> repository`.
- `/test` возвращает ровно:

```text
Hello!
```

- Для `/test` разрешён только `GET`; другие методы возвращают `405 Method Not Allowed`.
- Сохранён graceful shutdown по `SIGINT` / `SIGTERM`.

## Что добавлено в ЛР2

- PostgreSQL запускается в Docker.
- Go-сервер тоже запускается в Docker.
- `docker-compose.yml` поднимает два контейнера: `mini-avito-app` и `mini-avito-postgres`.
- Приложение подключается к БД внутри Docker-сети по адресу `postgres:5432`.
- Таблицы создаются автоматически при старте сервера на слое `repository`.
- `POST /dbtest` записывает строку из тела запроса в таблицу `db_test` и возвращает созданную запись в JSON.

## Запуск всей ЛР2 одной командой

```bash
docker compose up --build
```

Если нужно запустить в фоне:

```bash
docker compose up --build -d
```

Проверить контейнеры:

```bash
docker compose ps
```

Должны быть запущены:

```text
mini-avito-app
mini-avito-postgres
```

## Адреса

Go-сервер работает на порту `8080`:

```text
http://localhost:8080
```

PostgreSQL работает на порту `5432`, но его не нужно открывать в браузере.

В GitHub Codespaces открывать нужно ссылку именно с `8080`, например:

```text
https://...-8080.app.github.dev/test
```

Не открывать `5432` в браузере — это порт базы данных.

## Проверка ЛР1

```bash
curl -i http://localhost:8080/test
```

Ожидаемый ответ:

```text
HTTP/1.1 200 OK

Hello!
```

Проверка ограничения метода:

```bash
curl -i -X POST http://localhost:8080/test -d "abc"
```

Ожидаемый ответ:

```text
HTTP/1.1 405 Method Not Allowed

method not allowed
```

## Проверка ЛР2

```bash
curl -i -X POST http://localhost:8080/dbtest -d "hello database"
```

Ожидаемый ответ:

```text
HTTP/1.1 201 Created
Content-Type: application/json
```

И JSON примерно такого вида:

```json
{"id":1,"body":"hello database","created_at":"..."}
```

## Показать, что запись реально лежит в PostgreSQL

```bash
docker exec -it mini-avito-postgres psql -U postgres -d mini_avito -c "SELECT * FROM db_test ORDER BY id DESC;"
```

Там должна быть строка, которую отправляли через `/dbtest`.

## Остановить проект

```bash
docker compose down
```

Если нужно удалить ещё и данные БД:

```bash
docker compose down -v
```

## Что говорить на защите

В ЛР1 был HTTP-сервер с `/test`, чистой архитектурой и graceful shutdown. В ЛР2 мы добавили PostgreSQL, инициализацию таблиц в repository и POST-хэндлер `/dbtest`, который получает строку из тела запроса и сохраняет её в базу. Теперь всё запускается через Docker Compose: отдельный контейнер для Go-сервера и отдельный контейнер для PostgreSQL.
