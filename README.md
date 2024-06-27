### Для запуска
```shell
docker compose up
```

# Notes Service API

API для управления заметками пользователей.

## Роуты

- **GET /notes**
  - Описание: Получение списка всех заметок.
  - Headers: Не требуется аутентификация для публичных заметок либо Authorization: Bearer {token} для приватных заметок.

- **GET /notes/{id}**
  - Описание: Получение заметки по ID.
  - Headers: Не требуется аутентификация либо Authorization: Bearer {token} для приватных заметок.

- **POST /notes**
  - Описание: Создание новой заметки.
  - Headers:
    - Authorization: Bearer {token}

- **PATCH /notes/{id}**
  - Описание: Обновление заметки по ID.
  - Headers:
    - Authorization: Bearer {token}

- **DELETE /notes/{id}**
  - Описание: Удаление заметки по ID.
  - Headers:
    - Authorization: Bearer {token}


# Auth Service API

API для управления учетными записями пользователей.

## Роуты

- **POST /register**
  - Описание: Регистрация нового пользователя.

- **POST /login**
  - Описание: Аутентификация пользователя и выдача токена доступа.

- **GET /me**
  - Описание: Получение информации об аутентифицированном пользователе.
  - Headers:
    - Authorization: Bearer {token}

- **PATCH /me**
  - Описание: Изменение имени аутентифицированного пользователя.
  - Headers:
    - Authorization: Bearer {token}

- **PATCH /me/password**
  - Описание: Изменение пароля аутентифицированного пользователя.
  - Headers:
    - Authorization: Bearer {token}