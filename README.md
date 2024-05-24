<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->


- [Тестовое задание Go](#%D0%A2%D0%B5%D1%81%D1%82%D0%BE%D0%B2%D0%BE%D0%B5-%D0%B7%D0%B0%D0%B4%D0%B0%D0%BD%D0%B8%D0%B5-go)
  - [Формулировка задания](#%D0%A4%D0%BE%D1%80%D0%BC%D1%83%D0%BB%D0%B8%D1%80%D0%BE%D0%B2%D0%BA%D0%B0-%D0%B7%D0%B0%D0%B4%D0%B0%D0%BD%D0%B8%D1%8F)
  - [Описание решения](#%D0%9E%D0%BF%D0%B8%D1%81%D0%B0%D0%BD%D0%B8%D0%B5-%D1%80%D0%B5%D1%88%D0%B5%D0%BD%D0%B8%D1%8F)
  - [See Also](#see-also)
  - [User Guide](#user-guide)
    - [Полезные директории и файлы](#%D0%9F%D0%BE%D0%BB%D0%B5%D0%B7%D0%BD%D1%8B%D0%B5-%D0%B4%D0%B8%D1%80%D0%B5%D0%BA%D1%82%D0%BE%D1%80%D0%B8%D0%B8-%D0%B8-%D1%84%D0%B0%D0%B9%D0%BB%D1%8B)
    - [Описание основных команд для запуска](#%D0%9E%D0%BF%D0%B8%D1%81%D0%B0%D0%BD%D0%B8%D0%B5-%D0%BE%D1%81%D0%BD%D0%BE%D0%B2%D0%BD%D1%8B%D1%85-%D0%BA%D0%BE%D0%BC%D0%B0%D0%BD%D0%B4-%D0%B4%D0%BB%D1%8F-%D0%B7%D0%B0%D0%BF%D1%83%D1%81%D0%BA%D0%B0)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# Тестовое задание Go
______
## Формулировка задания
Реализовать систему для добавления и чтения постов и комментариев с использованием `GraphQL`, аналогичную комментариям к постам на популярных платформах, таких как `Хабр` или `Reddit`.

Характеристики системы постов:
- Можно просмотреть список постов.
- Можно просмотреть пост и комментарии под ним.
- Пользователь, написавший пост, может запретить оставление комментариев к своему посту.

Характеристики системы комментариев к постам:
-	Комментарии организованы иерархически, позволяя вложенность без ограничений.
-	Длина текста комментария ограничена до, например, 2000 символов.
-	Система пагинации для получения списка комментариев.

## Описание решения
Решение представляет из себя программу на `Go` с возможностью сборки/поставки в `Dockerfile` и запуском (например) через `docker-compose`.
**Основной ендпойнт GraphQL сервиса** - `/graphql` (если не менялись настройки то сервер работает на `localhost:8080`)

В качестве основного паттерна выбрана _layered_ архитектура, так как было важно иметь несколько реализаций для уровня доступа к данным. Также, данный паттерн
иммеет много плюсов для тестирования и поддержки кода (но к сожалению не в количестве кода на начальных этапах)

В своем решении я использовал сторонние библиотеки для генерации `GraphQL` ресолверов
и моделей ([gqlgen](https://gqlgen.com/)), для даталоадеров ([dataloaden](https://github.com/vikstrous/dataloadgen))

До этого я не был так тесно знаком (как теперь))) с `GraphQL` поэтому выбрал данные библиотеки, так как они предоставляют
строгую типизацию и имеют полезные втроенные фичи, также видел их использование в реальных продуктах

Если говорить про **решения типичных для GraphQL проблем**, то это
- **dataloader'ы** для смягчения **проблемы N+1**
- глубокая **вложенность запросов** контролируется используемой библиотекой [gqlgen](https://gqlgen.com/)

Основные файлы со **схемами GraphQL** можно найти в папке [/api](./api/graphql/). Как я уже говорил, до этого не использовал
GraphQL в своих проектах, поэтому не знаком со всеми best-practices в проектировании API для этой спецификации, но постарался в короткие сроки изучить
основные подходы и решения

**Систему пагинации** для постов и комментариев реализовал через [курсоры](https://relay.dev/graphql/connections.htm)

К сожалению, покрыть код unit-тестами не успел(

но на моих готовых проектах можно посмотреть, что я знаком с различными подходами тестирования)

## See Also
- [Jira-Analyzer](https://github.com/Jira-Analyzer/backend-services), Go
- [Telegram-Notifier](https://github.com/PonomarevAlexxander/telegram-notifier), Java
- [Swordfish Emulator](https://gitlab.com/IgorNikiforov/swordfish-emulator-go), Go

и другие репозитории в профиле)

## User Guide

### Полезные директории и файлы

Cхемы GraphQL можно найти в папке [/api](./api/graphql/)

Основной конфигурационный [файл](config/app/config.yaml) имеет в себе настройки для:
- Подключения к базе данных
- Выбора базы данных (in-memory/PostgreSQL)
- Настройки обычного сервера Go
- Настройки параметров GraphQL сервиса, а именно:
  - Максимальная сложность (max-complexity) - для предотвращения проблемы запросов с большой вложенностью.
  Позволяет чувствительно настраивать ограничения под нужны и конкретную нагрузку/требования
  - Настройки для даталоадеров:
    - Время ожидания до отправки батча
    - Максимальный размер батча (приоритетней времени ожидания)
- Настройки уровня логгирования

Миграционные файлы лежат в директории [/migrations](./migrations/)

docker-compose файлы для поднятия PostgreSQL и самого сервиса можно найти в папке [/deployments/docker](./deployments/docker/)

Multi-stage [Dockerfile](./Dockerfile) для сборки сервиса находится в корневой директории

### Описание основных команд для запуска
Загрузить локально исходные файлы репозитория можно через команду
```shell
git clone https://github.com/PonomarevAlexxander/graphql-forum.git
```

Основные шаги для сборки и запуска программы локально можно найти в [Makefile'e](Makefile):

- Сборка приложения осуществляется с помощью
```shell
make build
```

- Сборка `Docker-image` может быть запущена с помощью
```shell
make docker-build
```

- Загрузка утилит для кодогенерации и миграций (goose)
```shell
make install-tools
```

- Запуск миграций
```shell
make migration
```

- Кодогенерация для GraphQL
```shell
make gql-gen
```

- Запуск генераций go generate (в том числе даталоадеров)
```shell
make gogen
```

- Поднятие контейнера приложения и PosgreSQL в `docker-compose`
```shell
make start-dev
```

- Остановка контейнеров приложения и PostgreSQL
```shell
make stop-dev
```

- Запуск `Docker-image` приложения (`docker run`)
> Note: Через параметр `CONFIG_PATH` можно передать путь к yaml конфигу
```shell
make docker-run
```

- Запуск unit тестов
```shell
make unit-test
```
