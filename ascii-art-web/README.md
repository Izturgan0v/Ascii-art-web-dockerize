ascii-art-web — краткое руководство

О проекте
- Веб‑сервер на Go, генерирующий ASCII‑арт из введённого текста по выбранному баннеру.
- Слушает порт `:45674` (задан в `main.go`).
- Шаблоны HTML — в `templates/`, баннеры — в `ascii-art/banner/`.

Структура
- `main.go` — точка входа; маршруты (`/`, `/sabik`, `/ascii-art`); чтение шаблонов; обработка ошибок.
- `go.mod` — модуль `ascii-art-web`.
- `templates/` — `index.html`, `asciiart.html`, `error.html`, `sabik.html`.
- `ascii-art/` — генерация (`generator.go`), баннеры (`banner/*.txt`), утилиты (`utils/*`).

Запуск
- Выполнить: `go run ./main.go`
- Открыть: `http://localhost:45674/`

Маршруты
- `GET /` — главная (строго без query‑параметров). Любые query у корня → ошибка 400.
- `GET /sabik` — отдаёт `sabik.html`.
- `POST /ascii-art` — форма `application/x-www-form-urlencoded` с полями:
  - `text` — исходный текст
  - `banner` — один из: `standard`, `shadow`, `thinkertoy`
  Возвращает результат в `asciiart.html`.

Правила ввода
- Допустимые символы: ASCII 32–126, плюс переводы строк  LF (10).
- Последовательность `\n` интерпретируется как перенос строки.
- Пустые строки сохраняются в выводе.

Ошибки
- `400 Bad Request` — у `/` есть любые query‑параметры; неизвестный путь; баннер не из `standard|shadow|thinkertoy`; в тексте недопустимые символы.
- `405 Method Not Allowed` — неверный метод (например, не `POST` для `/ascii-art` или не `GET` для `/`).
- `500 Internal Server Error` — ошибка чтения шаблонов; отсутствует/повреждён баннер‑файл (проверяется по SHA‑256 в `utils`); внутренняя ошибка генерации.

Примеры
- Главная без query: `/` (OK). С query: `/?a=1` → 400.
- Генерация через curl:
  - `curl -X POST -d "text=Hello%0AWorld&banner=standard" http://localhost:45674/ascii-art`

Примечание
- Порт зашит в коде (`http.ListenAndServe(":45674", nil)`). Для изменения — поменяйте значение в `main.go` и перезапустите.


