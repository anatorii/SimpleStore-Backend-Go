FROM golang:1.24-alpine

WORKDIR /app

# Копируем файлы с зависимостями
COPY go.mod go.sum Makefile ./
RUN apk add --no-cache make
RUN make deps

# Копируем остальной код и собираем
COPY . ./
RUN make build

EXPOSE 8080
CMD ["./shop"]
