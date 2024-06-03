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
