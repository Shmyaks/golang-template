## Основная информация

### Реализовано в этом репозитории:
  - Реализация примера Echo - https://github.com/labstack/echo
    - Создание Routes
    - Создание Обработчиков 
  - Реализация примера Swagger - https://github.com/swaggo/swag
  - Реализция примера GORM -  https://gorm.io/docs/
    - Создание таблиц
    - Миграции   
  - Реализация примера JWT - https://github.com/golang-jwt/jwt
  - Реализация Docker с Postgres
  - Реализация PgAdmin

### Запуск:
1. Склонировать репозиторий
2. Настройка env (*см. далее*)
3. `docker-compose up -d`
4. Написать команду `go mod download`
5. Запустить `go run cmd/main.go`

### Env:
1. Редактируем .env *(Файл .env.example удаляем)*
```sh
DATABASE_HOST=
DATABASE_NAME=
DATABASE_USER=
DATABASE_PASSWORD=
DATABASE_PORT=5432

ADMIN_DEFAULT_EMAIL=
ADMIN_DEFAULT_PASSWORD=

JWT_SECRET=
JWT_ACCESS_TOKEN_EXPIRE_MINUTES=
```

### Документация:
`localhost:*ваш порт*/docs`
Чтобы обновить документацию - `swag init -g cmd/main.go`
