# Avito Test Task Autumn 2024

Данный проект является тестовым заданием и представляет из себя реализацию сервиса для создания тендеров и предложений.
В проекте реализованы ручки для создания, получения и изменения тендеров и предложений.

1. [Использованные технологии](#Использованные-технологии)
2. [Запуск приложения](#Запуск-приложения)
3. [Реализованные ручки](#реализованные-ручки)
4. [Дополнения](#дополнения)

## Использованные технологии
- Golang (go1.22.3 windows/amd64)
  - Доп. библиотеки:
    - github.com/jackc/pgx
    - github.com/Masterminds/squirrel
- openapitools/openapi-generator-cli
- Docker
- Postgres

`Pgx` и `squirrel` использовались для работы с базой данных. 
`openapi-generator-cli` был использован для создания структуры и базы проекта.
`Docker` был использован для сборки проекта. 

## Запуск приложения
Для начала необходимо склонировать репозиторий
```shell
git clone https://github.com/olga-sinepalnikova/avito-task-2024
```
Затем, перейдя в директорию проекта, выполнить команду:
```shell
make docker-up
```
> Предварительно нужно запустить докер!

Как только контейнер будет запущен, можно перейти по адресу `http://localhost:8080/api/ping`.
Если всё прошлоо успешно, то вы увидите: 'ok'.
Если же нет, то советую проверить занятые порты и при необходимости удалить/остановить последние контейнеры,
images и volumes.

Для остановки и удаления контейнера можно использовать ``Ctrl+C``, а затем
```shell
make docker-down
```

## Реализованные ручки

| Название группы    | Ручки                                  |
| ------------------ | -------------------------------------- |
| ping            | - /ping                                |
| tenders/new     | - /tenders/new                         |
| tenders/list    | - /tenders<br>- /tenders/my            |
| tenders/status  | - /tenders/status                      |
| tenders/version | - /tenders/edit<br>- /tenders/rollback |
| bids/new        | - /bids/new                            |
| bids/decision   | - /bids/submit_decision                |
| bids/list       | - /bids/list<br>- /bids/my             |
| bids/status     | - /bids/status                         |
| bids/version    | - /bids/edit<br>- /bids/rollback       |
| bids/feedback   | - /bids/reviews<br>- /bids/feedback    |


## Дополнения
Базы данных уже заполнены небольшим количеством информации, поэтому сразу можно попробовать вывести список тендеров или предложений.