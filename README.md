# Todo-App API

[![My Skills](https://skillicons.dev/icons?i=go,postgres,redis,kafka,docker)](https://skillicons.dev)

Многопользовательский todo-list на Go с JWT, Redis-кэшем и Kafka audit-log.

## Стек

**Backend:** Go 1.26, PostgreSQL, Redis, Apache Kafka (KRaft)

**Auth:** JWT + bcrypt

**Infra:** Docker Compose, golang-migrate

**Docs:** Swagger

**Logging:** zap

**Tests:** testify

## Возможности

- JWT-авторизация + blacklist токенов в Redis
- CRUD задач с изоляцией по пользователям
- Кэширование задач в Redis с инвалидацией (Cache-Aside)
- Audit-log через Kafka: события пишутся в отдельную таблицу в базе данных
- Swagger UI, структурированные логи (zap), unit-тесты

## Архитектура

Clean Architecture, слои `transport -> service -> repository`:

```
internal/
├── core/         # общее ядро (logger, errors, kafka, redis, jwt)
└── features/     # бизнес-фичи
    ├── users/    # register, login, me, logout
    ├── tasks/    # CRUD задач
    └── audit/    # запись событий в БД
```

## Запуск

```bash
cp .env.example .env    # заполнить секреты
make docker-build       # сборка образа + запуск
make docker-down        # остановить и удалить
```

- API: http://localhost:5678
- Swagger: http://localhost:5678/swagger/index.html
- Kafka UI: http://localhost:5679

## Тесты

```bash
go test -race ./...
```

## Ключевые решения

- **Cache-Aside** - инвалидация при мутациях, TTL как страховка
- **Best effort** для кэша и Kafka - их падение не роняет основной сценарий
- **Clean Architecture** - фичи изолированы, интерфейсы у потребителя, домен не знает про БД и Kafka

## API

| Метод | Путь | Описание |
|---|---|---|
| POST | /users/register | Регистрация |
| POST | /users/login | Логин, возвращает JWT |
| GET | /users/me | Текущий пользователь |
| PATCH | /users/me | Обновить профиль |
| DELETE | /users/me | Удалить профиль |
| POST | /users/logout | Выход (blacklist) |
| POST | /tasks | Создать задачу |
| GET | /tasks | Список задач |
| GET | /tasks/{id} | Получить задачу |
| PATCH | /tasks/{id} | Обновить задачу |
| DELETE | /tasks/{id} | Удалить задачу |

Все, кроме register/login, требуют `Authorization: Bearer <token>`.
