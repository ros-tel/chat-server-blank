# chat-server-blank
Заготовка учебного проекта

## Запуск под Windows x64
`bin\chat-server.exe`

## Запуск под Linux x64
`bin/chat-server`

## Документация к API
Локальная [swagger](http://localhost:8080/swagger/) позволит сделать запросы к API

GitHub [swagger](https://ros-tel.github.io/chat-server-blank/swagger/index.html)

## Для авторизации

Логины: `alexander`, `alisa`, `vitaly`, `vladimir`, `lada`, `semen`

Пароль у всех одинаковый `pass`

## Чат-сервер
Для тестирования сообщений в чате можно использовать расширение для Chrome [Smart Websocket Client](https://chrome.google.com/webstore/detail/smart-websocket-client/omalebghpgejjiaoknljcfmglgbpocdp)

В строке адреса вбиваете `ws://127.0.0.1:8080/ws/` и нажимаете кнопку `Connect`

Первым сообщением должно быть сообщение авторизации
Например:
```json
{
  "type": "auth",
  "auth": {
    "login": "alisa",
    "password": "pass"
  }
}
```

Все последующие сообщения в чата
В личку:
```json
{
  "type": "chat",
  "id": "qrta35",
  "to": 1,
  "chat": {
    "type": "text",
    "text": "Hello"
  }
}
```

В группу:
```json
{
  "type": "groupchat",
  "id": "qwrta35",
  "to": 1,
  "chat": {
    "type": "text",
    "text": "Hello"
  }
}
```

В ответ должны получать Эхо-сообщения с проставленным тегом `from` c выщим ID-пользователя
