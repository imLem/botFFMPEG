# Discord bot with ffmpeg

## Описание:
Так как в дискорде нет одинаковой поддержки видео форматов на всех устройствах, а также просто отсутствие поддержки некоторых кодеков, был сделан бот, который это фиксит.
- Hevc
ТикТок начал использовать кодек hevc в хайрез видео, а такой кодек не поддерживается дискордом, бот детектит видео с этим кодеком и конвертирует его в h.264, чтобы можно было смотреть его из приложения.
- Webm
Webm файлы нельзя просмотреть на устройствах apple, поэтому бот их конвертирует в mp4


## Запуск

!ВАЖНО! на вашей системе должен быть установлен ffmpeg

Из папки выполните команду, чтобы скомпилировать бота:
```
go build
```
Для запуска бота введите команду с вашим токеном от бота:
```
./botFFMPEG -t YOUR_BOT_TOKEN
```

## TODO

- [ ] Поддержка слеш команд
- [ ] Команда на смену языка(ru/eng)
- [ ] Добавление тестов
