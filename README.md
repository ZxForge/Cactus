# 🌵Cactus
>🚧 Находится в стадии разработки

Cactus - сервис рассылки уведомлений (почта, смс, push, webhook, sse, ws), с пользовательским интерфейсом и API

Для запуска миграций на локальной машине, нужно поднять докер и в корне проекта выполнить инструкцию:
```
dbmate --env-file ".env.dbmate.local" up 
```

Данная команда создаст полностью пустую структуру базы данных. 
> До версии v 1.0.0 нужно учитывать, что миграции не будут дополняться и нужна будет каждый раз удалять таблицы для обновления структуры.
 
Для запуска сидов для заполнения базы выполнить команду.
```
go run .\cmd\seeding\seeding.go
```
