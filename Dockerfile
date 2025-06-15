# Используем официальный образ Go как базовый
FROM golang:1.23-alpine

# Устанавливаем необходимые утилиты
RUN apk add --no-cache git

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем go.mod и go.sum
COPY go.mod go.sum ./
RUN go mod download

# Копируем все остальные файлы
COPY . .

# Собираем бинарник
RUN go build -o main .

# Указываем команду запуска
CMD ["/app/main"]
