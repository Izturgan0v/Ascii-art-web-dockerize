# Docker инструкции для ASCII Art Web приложения

## Обзор

Этот проект использует Docker для контейнеризации Go веб-приложения, которое генерирует ASCII искусство.

## Структура Docker файлов

- `Dockerfile` - основной файл для сборки Docker образа
- `.dockerignore` - исключает ненужные файлы из Docker контекста
- `docker-compose.yml` - для удобного запуска и управления контейнером

## Быстрый старт

### 1. Сборка и запуск с Docker Compose (рекомендуется)

```bash
# Собрать образ и запустить контейнер
docker-compose up --build

# Запустить в фоновом режиме
docker-compose up -d --build

# Остановить контейнер
docker-compose down
```

### 2. Ручная сборка с Docker

```bash
# Собрать Docker образ
docker build -t ascii-art-web .

# Запустить контейнер
docker run -d -p 45674:45674 --name ascii-art-web-container ascii-art-web

# Остановить и удалить контейнер
docker stop ascii-art-web-container
docker rm ascii-art-web-container
```

## Доступ к приложению

После запуска приложение будет доступно по адресу:
- **Локально**: http://localhost:45674
- **В сети**: http://[IP-адрес-вашего-компьютера]:45674

## Управление контейнером

### Просмотр логов

```bash
# Логи в реальном времени
docker-compose logs -f

# Последние 100 строк логов
docker-compose logs --tail=100
```

### Проверка состояния

```bash
# Статус контейнеров
docker-compose ps

# Статистика использования ресурсов
docker stats ascii-art-web-container
```

### Перезапуск

```bash
# Перезапустить сервис
docker-compose restart

# Пересобрать и перезапустить
docker-compose up --build -d
```

## Разработка

### Режим разработки с горячей перезагрузкой

Для разработки можно раскомментировать volumes в `docker-compose.yml`:

```yaml
volumes:
  - ./templates:/app/templates:ro
  - ./ascii-art:/app/ascii-art:ro
```

Это позволит изменять шаблоны и баннеры без пересборки образа.

### Отладка

```bash
# Запуск в интерактивном режиме
docker-compose run --rm ascii-art-web sh

# Просмотр файлов внутри контейнера
docker exec -it ascii-art-web-container sh
```

## Оптимизация

### Многоэтапная сборка

Dockerfile использует многоэтапную сборку:
1. **Этап builder**: компилирует Go приложение
2. **Этап runtime**: создает минимальный образ для запуска

Это уменьшает размер финального образа.

### Кэширование слоев

- Зависимости Go загружаются отдельно для лучшего кэширования
- Исходный код копируется после загрузки зависимостей

## Безопасность

- Приложение запускается от непривилегированного пользователя
- Используется минимальный базовый образ (Alpine Linux)
- Отключен CGO для статической компиляции

## Мониторинг

### Health Check

Контейнер включает проверку здоровья, которая:
- Проверяет доступность приложения каждые 30 секунд
- Позволяет Docker определить, когда контейнер готов к работе

### Логирование

- Логи ограничены размером 10MB
- Хранится максимум 3 файла логов
- Автоматическая ротация логов

## Troubleshooting

### Проблемы с портами

```bash
# Проверить, какие порты заняты
netstat -tulpn | grep 45674

# Остановить процесс, занимающий порт
sudo lsof -ti:45674 | xargs kill -9
```

### Проблемы с правами доступа

```bash
# Исправить права доступа к файлам
sudo chown -R $USER:$USER .

# Дать права на выполнение Docker
sudo usermod -aG docker $USER
```

### Очистка Docker

```bash
# Удалить неиспользуемые образы
docker image prune -a

# Удалить неиспользуемые контейнеры
docker container prune

# Полная очистка
docker system prune -a
```

## Производительность

### Ограничения ресурсов

В `docker-compose.yml` настроены ограничения:
- **Память**: максимум 512MB, минимум 256MB
- **CPU**: максимум 0.5 ядра, минимум 0.25 ядра

### Масштабирование

```bash
# Запустить несколько экземпляров
docker-compose up --scale ascii-art-web=3
```

## Сеть

### Пользовательские сети

```bash
# Создать пользовательскую сеть
docker network create ascii-art-network

# Подключить контейнер к сети
docker run --network ascii-art-network ascii-art-web
```

## Обновления

### Обновление образа

```bash
# Получить последние изменения
git pull

# Пересобрать образ
docker-compose build --no-cache

# Перезапустить с новым образом
docker-compose up -d
```

## Резервное копирование

### Сохранение данных

```bash
# Создать резервную копию контейнера
docker commit ascii-art-web-container ascii-art-web:backup

# Сохранить образ в файл
docker save ascii-art-web:backup > ascii-art-backup.tar

# Загрузить образ из файла
docker load < ascii-art-backup.tar
``` 