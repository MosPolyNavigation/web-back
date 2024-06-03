# web-back
## Запуск программы на сервере
1. Создать файл .env см. .env-example
2. Скачать Docker
3. Выполнить комманду в терминале
   ```shell
   docker compose up -d --build
   ```
5. Программа будет доступна по адресу http://HOST:SERVER_PORT/
   Где HOST - адресс сервера, а SERVER_PORT - переменная указанная в .env
## Адресс сайта 
На данный момент сайт доступ к серверу осуществляется по адрессу https://navigation.nzb3.su/
## Запросы
Пример запроса на получение плана
```
https://navigation.nzb3.su/plan?campus=bs&corpus=n&floor=1
```
